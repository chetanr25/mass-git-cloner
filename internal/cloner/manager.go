package cloner

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/chetanr25/mass-git-cloner/internal/config"
	"github.com/chetanr25/mass-git-cloner/internal/ui"
	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

// Manager manages the overall cloning process
type Manager struct {
	config     *config.Config
	workerPool *WorkerPool
	progress   *ui.ProgressTracker
}

// NewManager creates a new cloning manager
func NewManager(cfg *config.Config) *Manager {
	return &Manager{
		config:     cfg,
		workerPool: NewWorkerPool(cfg),
	}
}

// CloneRepositories clones multiple repositories concurrently
func (m *Manager) CloneRepositories(repos []*models.Repository, username string) error {
	if len(repos) == 0 {
		return fmt.Errorf("no repositories to clone")
	}

	// Check if git is available
	if err := CheckGitInstalled(); err != nil {
		return err
	}

	// Prepare target directory
	targetDir, err := PrepareTargetDirectory(username, m.config.BaseDir)
	if err != nil {
		return fmt.Errorf("failed to prepare target directory: %w", err)
	}

	ui.DisplayInfo(fmt.Sprintf("Cloning %d repositories to: %s", len(repos), targetDir))

	// Set up context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		ui.DisplayInfo("\nReceived interrupt signal. Cancelling remaining clones...")
		cancel()
	}()

	// Initialize progress tracker
	m.progress = ui.NewProgressTracker(len(repos))

	// Start worker pool
	m.workerPool.Start(ctx)

	// Process results in a separate goroutine
	resultsDone := make(chan bool)
	go m.processResults(resultsDone)

	// Submit jobs
	for _, repo := range repos {
		job := &CloneJob{
			Repository: repo,
			TargetDir:  targetDir,
		}
		
		select {
		case <-ctx.Done():
			break
		default:
			m.workerPool.AddJob(job)
		}
	}

	// Wait for all results to be processed
	<-resultsDone

	// Stop worker pool
	m.workerPool.Stop()

	// Finish progress tracking
	m.progress.Finish()

	return nil
}

// processResults processes cloning results and updates progress
func (m *Manager) processResults(done chan<- bool) {
	completed := 0
	total := m.progress.Total()

	for result := range m.workerPool.Results() {
		if result.Success {
			m.progress.Success(result.Repository.Name)
		} else {
			m.progress.Failure(result.Repository.Name, result.Error)
		}

		completed++
		if completed >= total {
			break
		}
	}

	done <- true
}

// CloneRepositoriesWithRetry clones repositories with retry logic
func (m *Manager) CloneRepositoriesWithRetry(repos []*models.Repository, username string) error {
	maxRetries := m.config.RetryAttempts
	var failedRepos []*models.Repository

	for attempt := 1; attempt <= maxRetries; attempt++ {
		reposToClone := repos
		if attempt > 1 {
			reposToClone = failedRepos
			if len(reposToClone) == 0 {
				break
			}
			ui.DisplayInfo(fmt.Sprintf("Retrying failed repositories (attempt %d/%d)", attempt, maxRetries))
		}

		err := m.CloneRepositories(reposToClone, username)
		if err != nil {
			return err
		}

		// Collect failed repositories for retry
		if attempt < maxRetries {
			failedRepos = m.getFailedRepositories(reposToClone, username)
		}
	}

	return nil
}

// getFailedRepositories returns repositories that failed to clone
func (m *Manager) getFailedRepositories(repos []*models.Repository, username string) []*models.Repository {
	var failed []*models.Repository
	targetDir, _ := PrepareTargetDirectory(username, m.config.BaseDir)

	for _, repo := range repos {
		info, err := GetRepositoryInfo(fmt.Sprintf("%s/%s", targetDir, repo.Name))
		if err != nil || !info.Exists || !info.IsGitRepo {
			failed = append(failed, repo)
		}
	}

	return failed
}

// UpdateRepositories updates existing repositories
func (m *Manager) UpdateRepositories(repos []*models.Repository, username string) error {
	targetDir, err := PrepareTargetDirectory(username, m.config.BaseDir)
	if err != nil {
		return err
	}

	ctx := context.Background()
	cloner := NewGitCloner(m.config)
	progress := ui.NewProgressTracker(len(repos))

	for _, repo := range repos {
		repoPath := fmt.Sprintf("%s/%s", targetDir, repo.Name)
		progress.Update(fmt.Sprintf("Updating %s...", repo.Name))

		if err := cloner.UpdateRepository(ctx, repoPath); err != nil {
			progress.Failure(repo.Name, err)
		} else {
			progress.Success(repo.Name)
		}
	}

	progress.Finish()
	return nil
}