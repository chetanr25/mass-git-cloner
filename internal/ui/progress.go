package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

// ProgressTracker tracks and displays cloning progress
type ProgressTracker struct {
	total     int
	completed int
	failed    int
	current   string
	startTime time.Time
}

// NewProgressTracker creates a new progress tracker
func NewProgressTracker(total int) *ProgressTracker {
	return &ProgressTracker{
		total:     total,
		completed: 0,
		failed:    0,
		current:   "",
		startTime: time.Now(),
	}
}

// Update updates the progress with the current operation
func (p *ProgressTracker) Update(current string) {
	p.current = current
	p.display()
}

// Success marks a repository as successfully cloned
func (p *ProgressTracker) Success(repoName string) {
	p.completed++
	p.current = fmt.Sprintf("✅ %s", repoName)
	p.display()
	time.Sleep(100 * time.Millisecond) // Brief pause to show the success
}

// Failure marks a repository as failed to clone
func (p *ProgressTracker) Failure(repoName string, err error) {
	p.failed++
	p.current = fmt.Sprintf("❌ %s: %v", repoName, err)
	p.display()
	time.Sleep(500 * time.Millisecond) // Longer pause to show the error
}

// display shows the current progress
func (p *ProgressTracker) display() {
	// Move cursor to beginning of line and clear it
	fmt.Print("\r" + strings.Repeat(" ", 100) + "\r")

	percentage := float64(p.completed+p.failed) / float64(p.total) * 100
	elapsed := time.Since(p.startTime)

	// Create progress bar
	barWidth := 30
	filled := int(float64(barWidth) * percentage / 100)
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

	status := fmt.Sprintf("[%s] %.1f%% (%d/%d) %s",
		bar,
		percentage,
		p.completed+p.failed,
		p.total,
		elapsed.Truncate(time.Second),
	)

	fmt.Print(status)

	if p.current != "" {
		fmt.Printf("\n%s", p.current)
		if p.completed+p.failed < p.total {
			fmt.Print("\n")
		}
	}
}

// Finish completes the progress tracking
func (p *ProgressTracker) Finish() {
	fmt.Print("\r" + strings.Repeat(" ", 100) + "\r")
	elapsed := time.Since(p.startTime)
	
	fmt.Printf("🎉 Cloning completed!\n")
	fmt.Printf("   Total: %d repositories\n", p.total)
	fmt.Printf("   Successful: %d\n", p.completed)
	fmt.Printf("   Failed: %d\n", p.failed)
	fmt.Printf("   Duration: %s\n", elapsed.Truncate(time.Second))
	
	if p.failed > 0 {
		fmt.Printf("   Success rate: %.1f%%\n", float64(p.completed)/float64(p.total)*100)
	}
}

// Total returns the total number of repositories
func (p *ProgressTracker) Total() int {
	return p.total
}

// DisplayStats shows repository statistics
func DisplayStats(stats *models.RepositoryStats) {
	fmt.Printf("\n📊 Repository Statistics\n")
	fmt.Printf("========================\n")
	fmt.Printf("Total repositories: %d\n", stats.Total)
	fmt.Printf("Non-fork repositories: %d\n", stats.NonForks)
	fmt.Printf("Fork repositories: %d\n", stats.Forks)
	fmt.Printf("Public repositories: %d\n", stats.Public)
	fmt.Printf("Private repositories: %d\n", stats.Private)
	fmt.Println()
}

// DisplayFilterMenu shows the filter selection menu
func DisplayFilterMenu(stats *models.RepositoryStats) models.FilterType {
	fmt.Printf("📂 Select repository type to clone:\n")
	fmt.Printf("1. All repositories (%d)\n", stats.Total)
	fmt.Printf("2. Non-fork repositories only (%d)\n", stats.NonForks)
	fmt.Printf("3. Fork repositories only (%d)\n", stats.Forks)
	fmt.Print("\nEnter your choice (1-3): ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 2:
		return models.FilterNonForks
	case 3:
		return models.FilterForksOnly
	default:
		return models.FilterAll
	}
}