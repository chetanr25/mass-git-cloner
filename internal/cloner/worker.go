package cloner

import (
	"context"
	"sync"
	"time"

	"github.com/chetanr25/mass-git-cloner/internal/config"
	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

// Worker represents a cloning worker
type Worker struct {
	id      int
	cloner  *GitCloner
	results chan<- *models.CloneResult
	jobs    <-chan *CloneJob
}

// CloneJob represents a cloning job
type CloneJob struct {
	Repository *models.Repository
	TargetDir  string
}

// NewWorker creates a new worker
func NewWorker(id int, cloner *GitCloner, jobs <-chan *CloneJob, results chan<- *models.CloneResult) *Worker {
	return &Worker{
		id:      id,
		cloner:  cloner,
		results: results,
		jobs:    jobs,
	}
}

// Start starts the worker
func (w *Worker) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-w.jobs:
			if !ok {
				return
			}
			w.processJob(ctx, job)
		case <-ctx.Done():
			return
		}
	}
}

// processJob processes a single cloning job
func (w *Worker) processJob(ctx context.Context, job *CloneJob) {
	start := time.Now()
	result := &models.CloneResult{
		Repository: job.Repository,
		Success:    false,
	}

	var err error
	err = w.cloner.CloneRepository(ctx, job.Repository, job.TargetDir)

	result.Duration = time.Since(start)
	result.Success = err == nil
	result.Error = err

	select {
	case w.results <- result:
	case <-ctx.Done():
		return
	}
}

// WorkerPool manages a pool of workers for concurrent cloning
type WorkerPool struct {
	config    *config.Config
	cloner    *GitCloner
	workers   []*Worker
	jobs      chan *CloneJob
	results   chan *models.CloneResult
	wg        sync.WaitGroup
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(cfg *config.Config) *WorkerPool {
	cloner := NewGitCloner(cfg)
	
	return &WorkerPool{
		config:  cfg,
		cloner:  cloner,
		jobs:    make(chan *CloneJob, cfg.MaxConcurrency*2),
		results: make(chan *models.CloneResult, cfg.MaxConcurrency*2),
	}
}

// Start starts the worker pool
func (wp *WorkerPool) Start(ctx context.Context) {
	// Create workers
	for i := 0; i < wp.config.MaxConcurrency; i++ {
		worker := NewWorker(i, wp.cloner, wp.jobs, wp.results)
		wp.workers = append(wp.workers, worker)
		
		wp.wg.Add(1)
		go worker.Start(ctx, &wp.wg)
	}
}

// AddJob adds a cloning job to the queue
func (wp *WorkerPool) AddJob(job *CloneJob) {
	wp.jobs <- job
}

// Results returns the results channel
func (wp *WorkerPool) Results() <-chan *models.CloneResult {
	return wp.results
}

// Stop stops the worker pool
func (wp *WorkerPool) Stop() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
}