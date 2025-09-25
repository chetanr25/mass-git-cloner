package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chetanr25/mass-git-cloner/internal/cloner"
	"github.com/chetanr25/mass-git-cloner/internal/config"
	"github.com/chetanr25/mass-git-cloner/internal/github"
	"github.com/chetanr25/mass-git-cloner/internal/ui"
)

func main() {
	// Display welcome message
	ui.DisplayWelcome()

	// Load configuration
	cfg := config.LoadConfig()

	// Create GitHub client
	client := github.NewClient(cfg)

	// Get username from user
	username, err := ui.PromptUsername()
	if err != nil {
		ui.DisplayError(err)
		os.Exit(1)
	}

	// Check if user exists
	ui.DisplayInfo(fmt.Sprintf("Checking if user '%s' exists...", username))
	exists, err := client.UserExists(username)
	if err != nil {
		ui.DisplayError(fmt.Errorf("failed to check user existence: %w", err))
		os.Exit(1)
	}

	if !exists {
		ui.DisplayError(fmt.Errorf("user or organization '%s' not found", username))
		os.Exit(1)
	}

	// Fetch repositories
	ui.DisplayInfo("Fetching repositories...")
	repos, err := client.GetRepositories(username)
	if err != nil {
		ui.DisplayError(fmt.Errorf("failed to fetch repositories: %w", err))
		os.Exit(1)
	}

	if len(repos) == 0 {
		ui.DisplayInfo("No repositories found for this user.")
		return
	}

	// Calculate statistics
	stats := github.CalculateStats(repos)

	// Display statistics with beautiful Bubbletea UI
	if err := ui.ShowRepositoryStats(stats, username); err != nil {
		ui.DisplayError(fmt.Errorf("failed to display statistics: %w", err))
		os.Exit(1)
	}

	// Let user choose filter type with Bubbletea
	filterType, err := ui.ShowFilterSelector(stats)
	if err != nil {
		ui.DisplayError(fmt.Errorf("filter selection failed: %w", err))
		os.Exit(1)
	}

	// Filter repositories
	filteredRepos := github.FilterRepositories(repos, filterType)

	if len(filteredRepos) == 0 {
		ui.DisplayInfo("No repositories match the selected filter.")
		return
	}

	// Interactive repository selection with beautiful Bubbletea UI
	selectedRepos, err := ui.ShowRepositorySelector(filteredRepos, filterType)
	if err != nil {
		ui.DisplayError(fmt.Errorf("repository selection failed: %w", err))
		os.Exit(1)
	}

	if len(selectedRepos) == 0 {
		ui.DisplayInfo("No repositories selected for cloning.")
		return
	}

	ui.DisplaySuccess(fmt.Sprintf("Selected %d repositories for cloning", len(selectedRepos)))

	// Start cloning process
	manager := cloner.NewManager(cfg)

	ui.DisplayInfo("Starting repository cloning...")
	if err := manager.CloneRepositoriesWithRetry(selectedRepos, username); err != nil {
		ui.DisplayError(fmt.Errorf("cloning failed: %w", err))
		os.Exit(1)
	}

	ui.DisplaySuccess("All operations completed!")
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
