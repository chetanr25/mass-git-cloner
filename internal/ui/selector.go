package ui

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

// Selector handles interactive repository selection
type Selector struct {
	repositories []*models.Repository
	selected     map[int]bool
	cursor       int
	filter       models.FilterType
}

// NewSelector creates a new repository selector
func NewSelector(repos []*models.Repository, filter models.FilterType) *Selector {
	selected := make(map[int]bool)
	return &Selector{
		repositories: repos,
		selected:     selected,
		cursor:       0,
		filter:       filter,
	}
}

// Show displays the interactive selector (basic version without external dependencies)
func (s *Selector) Show() ([]*models.Repository, error) {
	if len(s.repositories) == 0 {
		return nil, fmt.Errorf("no repositories to select from")
	}

	fmt.Printf("\n📋 Repository Selection (%s)\n", s.filter.String())
	fmt.Println("Instructions:")
	fmt.Println("- Enter repository numbers to select/deselect (e.g., 1,3,5 or 1-5)")
	fmt.Println("- Type 'all' to select all repositories")
	fmt.Println("- Type 'none' to deselect all repositories")
	fmt.Println("- Type 'done' when finished selecting")
	fmt.Println()

	s.displayRepositories()

	for {
		fmt.Print("\nEnter your selection: ")
		
		var input string
		fmt.Scanln(&input)
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "done" {
			break
		}

		if input == "all" {
			s.selectAll()
			s.displayRepositories()
			continue
		}

		if input == "none" {
			s.selectNone()
			s.displayRepositories()
			continue
		}

		if err := s.processSelection(input); err != nil {
			fmt.Printf("❌ Invalid selection: %v\n", err)
			continue
		}

		s.displayRepositories()
	}

	return s.getSelectedRepositories(), nil
}

// displayRepositories shows the current list of repositories with selection status
func (s *Selector) displayRepositories() {
	// Clear screen (basic cross-platform approach)
	fmt.Print("\033[H\033[2J")

	fmt.Printf("\n📋 Repository Selection (%s) - %d selected\n", s.filter.String(), len(s.selected))
	fmt.Println(strings.Repeat("=", 80))

	for i, repo := range s.repositories {
		status := "[ ]"
		if s.selected[i] {
			status = "[✓]"
		}

		language := repo.Language
		if language == "" {
			language = "N/A"
		}

		description := repo.Description
		if len(description) > 50 {
			description = description[:47] + "..."
		}
		if description == "" {
			description = "No description"
		}

		fmt.Printf("%s %2d. %-30s %-12s ⭐ %-4d %s\n",
			status,
			i+1,
			repo.Name,
			language,
			repo.StarCount,
			description,
		)
	}

	fmt.Println(strings.Repeat("=", 80))
}

// processSelection processes user input for repository selection
func (s *Selector) processSelection(input string) error {
	// Handle comma-separated values and ranges
	parts := strings.Split(input, ",")
	
	for _, part := range parts {
		part = strings.TrimSpace(part)
		
		// Handle ranges (e.g., "1-5")
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return fmt.Errorf("invalid range format: %s", part)
			}

			start, err := strconv.Atoi(strings.TrimSpace(rangeParts[0]))
			if err != nil {
				return fmt.Errorf("invalid start number in range: %s", rangeParts[0])
			}

			end, err := strconv.Atoi(strings.TrimSpace(rangeParts[1]))
			if err != nil {
				return fmt.Errorf("invalid end number in range: %s", rangeParts[1])
			}

			if start > end {
				start, end = end, start
			}

			for i := start; i <= end; i++ {
				if err := s.toggleSelection(i); err != nil {
					return err
				}
			}
		} else {
			// Handle single number
			num, err := strconv.Atoi(part)
			if err != nil {
				return fmt.Errorf("invalid number: %s", part)
			}

			if err := s.toggleSelection(num); err != nil {
				return err
			}
		}
	}

	return nil
}

// toggleSelection toggles the selection status of a repository
func (s *Selector) toggleSelection(num int) error {
	index := num - 1 // Convert to 0-based index
	
	if index < 0 || index >= len(s.repositories) {
		return fmt.Errorf("repository number %d is out of range (1-%d)", num, len(s.repositories))
	}

	if s.selected[index] {
		delete(s.selected, index)
	} else {
		s.selected[index] = true
	}

	return nil
}

// selectAll selects all repositories
func (s *Selector) selectAll() {
	s.selected = make(map[int]bool)
	for i := range s.repositories {
		s.selected[i] = true
	}
}

// selectNone deselects all repositories
func (s *Selector) selectNone() {
	s.selected = make(map[int]bool)
}

// getSelectedRepositories returns the list of selected repositories
func (s *Selector) getSelectedRepositories() []*models.Repository {
	selected := make([]*models.Repository, 0, len(s.selected))
	
	for index := range s.selected {
		if index < len(s.repositories) {
			selected = append(selected, s.repositories[index])
		}
	}

	return selected
}

// clearScreen clears the terminal screen
func clearScreen() {
	cmd := exec.Command("clear")
	if os.Getenv("OS") == "Windows_NT" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}