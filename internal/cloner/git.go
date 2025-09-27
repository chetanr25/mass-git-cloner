package cloner

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/chetanr25/mass-git-cloner/internal/config"
	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

type GitCloner struct {
	config *config.Config
}

func NewGitCloner(cfg *config.Config) *GitCloner {
	return &GitCloner{
		config: cfg,
	}
}

func (g *GitCloner) CloneRepository(ctx context.Context, repo *models.Repository, targetDir string) error {

	repoPath := filepath.Join(targetDir, repo.Name)

	if _, err := os.Stat(repoPath); err == nil {
		return fmt.Errorf("directory already exists: %s", repoPath)
	}

	cloneCtx, cancel := context.WithTimeout(ctx, g.config.CloneTimeout)
	defer cancel()

	cmd := exec.CommandContext(cloneCtx, "git", "clone", repo.CloneURL, repoPath)

	if err := cmd.Run(); err != nil {
		os.RemoveAll(repoPath)
		return fmt.Errorf("git clone failed: %w", err)
	}

	return nil
}

func (g *GitCloner) UpdateRepository(ctx context.Context, repoPath string) error {

	gitDir := filepath.Join(repoPath, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return fmt.Errorf("not a git repository: %s", repoPath)
	}

	updateCtx, cancel := context.WithTimeout(ctx, g.config.CloneTimeout)
	defer cancel()

	cmd := exec.CommandContext(updateCtx, "git", "-C", repoPath, "pull", "--ff-only")
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git pull failed: %w", err)
	}

	return nil
}

func CheckGitInstalled() error {
	cmd := exec.Command("git", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git is not installed or not available in PATH")
	}
	return nil
}

func PrepareTargetDirectory(username, baseDir string) (string, error) {
	targetDir := filepath.Join(baseDir, username)

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create target directory: %w", err)
	}

	return targetDir, nil
}

func GetRepositoryInfo(repoPath string) (*LocalRepoInfo, error) {
	info := &LocalRepoInfo{
		Path:   repoPath,
		Exists: false,
	}

	if stat, err := os.Stat(repoPath); err != nil {
		return info, nil
	} else if !stat.IsDir() {
		return info, nil
	}

	info.Exists = true

	gitDir := filepath.Join(repoPath, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		info.IsGitRepo = true

		if stat, err := os.Stat(repoPath); err == nil {
			info.LastModified = stat.ModTime()
		}
	}

	return info, nil
}

type LocalRepoInfo struct {
	Path         string
	Exists       bool
	IsGitRepo    bool
	LastModified time.Time
}
