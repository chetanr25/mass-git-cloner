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

// GitCloner handles git operations
type GitCloner struct {
	config *config.Config
}

// NewGitCloner creates a new git cloner
func NewGitCloner(cfg *config.Config) *GitCloner {
	return &GitCloner{
		config: cfg,
	}
}

// CloneRepository clones a single repository
func (g *GitCloner) CloneRepository(ctx context.Context, repo *models.Repository, targetDir string) error {
	// Create target directory if it doesn't exist
	repoPath := filepath.Join(targetDir, repo.Name)
	
	// Check if directory already exists
	if _, err := os.Stat(repoPath); err == nil {
		return fmt.Errorf("directory already exists: %s", repoPath)
	}

	// Create context with timeout
	cloneCtx, cancel := context.WithTimeout(ctx, g.config.CloneTimeout)
	defer cancel()

	// Prepare git clone command
	cmd := exec.CommandContext(cloneCtx, "git", "clone", repo.CloneURL, repoPath)
	
	// Set up environment
	cmd.Env = os.Environ()
	
	// Execute the command
	if err := cmd.Run(); err != nil {
		// Clean up partial clone on failure
		os.RemoveAll(repoPath)
		return fmt.Errorf("git clone failed: %w", err)
	}

	return nil
}

// UpdateRepository pulls the latest changes for an existing repository
func (g *GitCloner) UpdateRepository(ctx context.Context, repoPath string) error {
	// Check if it's a git repository
	gitDir := filepath.Join(repoPath, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		return fmt.Errorf("not a git repository: %s", repoPath)
	}

	updateCtx, cancel := context.WithTimeout(ctx, g.config.CloneTimeout)
	defer cancel()

	// Change to repository directory and pull
	cmd := exec.CommandContext(updateCtx, "git", "-C", repoPath, "pull", "--ff-only")
	cmd.Env = os.Environ()
	
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git pull failed: %w", err)
	}

	return nil
}

// CheckGitInstalled verifies that git is installed and available
func CheckGitInstalled() error {
	cmd := exec.Command("git", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git is not installed or not available in PATH")
	}
	return nil
}

// PrepareTargetDirectory creates the target directory structure
func PrepareTargetDirectory(username, baseDir string) (string, error) {
	targetDir := filepath.Join(baseDir, username)
	
	// Create directory if it doesn't exist
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create target directory: %w", err)
	}

	return targetDir, nil
}

// GetRepositoryInfo gets information about an existing local repository
func GetRepositoryInfo(repoPath string) (*LocalRepoInfo, error) {
	info := &LocalRepoInfo{
		Path:   repoPath,
		Exists: false,
	}

	// Check if directory exists
	if stat, err := os.Stat(repoPath); err != nil {
		return info, nil
	} else if !stat.IsDir() {
		return info, nil
	}

	info.Exists = true

	// Check if it's a git repository
	gitDir := filepath.Join(repoPath, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		info.IsGitRepo = true

		// Get last modified time
		if stat, err := os.Stat(repoPath); err == nil {
			info.LastModified = stat.ModTime()
		}
	}

	return info, nil
}

// LocalRepoInfo represents information about a local repository
type LocalRepoInfo struct {
	Path         string
	Exists       bool
	IsGitRepo    bool
	LastModified time.Time
}