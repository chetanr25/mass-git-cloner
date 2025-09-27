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

type Manager struct {
	config   *config.Config
	cloner   *GitCloner
	progress *ui.ProgressTracker
}

func NewManager(cfg *config.Config) *Manager {
	return &Manager{
		config: cfg,
		cloner: NewGitCloner(cfg),
	}
}

func (m *Manager) CloneRepositories(repos []*models.Repository, username string) error {
	if len(repos) == 0 {
		return fmt.Errorf("no repositories to clone")
	}

	if err := CheckGitInstalled(); err != nil {
		return err
	}

	targetDir, err := PrepareTargetDirectory(username, m.config.BaseDir)
	if err != nil {
		return fmt.Errorf("failed to prepare target directory: %w", err)
	}

	ui.DisplayInfo(fmt.Sprintf("Cloning %d repositories to: %s", len(repos), targetDir))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		ui.DisplayInfo("\nReceived interrupt signal. Stopping...")
		cancel()
	}()

	m.progress = ui.InitProgressTracker(len(repos))

	for i, repo := range repos {
		select {
		case <-ctx.Done():
			ui.DisplayInfo("Cloning stopped by user")
			return nil
		default:
		}

		m.progress.Update(fmt.Sprintf("Cloning %s (%d/%d)...", repo.Name, i+1, len(repos)))

		err := m.cloner.CloneRepository(ctx, repo, targetDir)

		if err != nil {
			m.progress.Failure(repo.Name, err)
		} else {
			m.progress.Success(repo.Name)
		}
	}

	m.progress.Finish()

	return nil
}

func (m *Manager) UpdateRepositories(repos []*models.Repository, username string) error {
	targetDir, err := PrepareTargetDirectory(username, m.config.BaseDir)
	if err != nil {
		return err
	}

	ctx := context.Background()
	progress := ui.InitProgressTracker(len(repos))

	for i, repo := range repos {
		repoPath := fmt.Sprintf("%s/%s", targetDir, repo.Name)
		progress.Update(fmt.Sprintf("Updating %s (%d/%d)...", repo.Name, i+1, len(repos)))

		if err := m.cloner.UpdateRepository(ctx, repoPath); err != nil {
			progress.Failure(repo.Name, err)
		} else {
			progress.Success(repo.Name)
		}
	}

	progress.Finish()
	return nil
}
