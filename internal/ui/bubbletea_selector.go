package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

// Styles for the UI
var (
	// Main container styles
	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7C3AED")).
		Padding(0, 1).
		Bold(true)

	headerStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#3B82F6")).
		Padding(0, 1)

	// Repository item styles
	selectedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#10B981")).
		Padding(0, 1).
		Bold(true)

	checkedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#10B981")).
		Bold(true)

	uncheckedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6B7280"))

	cursorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F59E0B")).
		Bold(true)

	// Language tag styles
	languageStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#8B5CF6")).
		Padding(0, 1).
		MarginRight(1)

	// Stats styles
	starsStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FBBF24"))

	forksStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#34D399"))

	// Footer styles
	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6B7280")).
		Margin(1, 0)

	confirmStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#EF4444")).
		Padding(0, 1).
		Bold(true)
)

// RepositorySelectorModel represents the Bubbletea model for repository selection
type RepositorySelectorModel struct {
	repositories []*models.Repository
	selected     map[int]bool
	cursor       int
	filter       models.FilterType
	showConfirm  bool
	confirmed    bool
	done         bool
	width        int
	height       int
}

// NewRepositorySelectorModel creates a new repository selector model
func NewRepositorySelectorModel(repos []*models.Repository, filter models.FilterType) *RepositorySelectorModel {
	return &RepositorySelectorModel{
		repositories: repos,
		selected:     make(map[int]bool),
		cursor:       0,
		filter:       filter,
		showConfirm:  false,
		confirmed:    false,
		done:         false,
		width:        80,
		height:       24,
	}
}

// Init initializes the model
func (m *RepositorySelectorModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m *RepositorySelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		if m.showConfirm {
			return m.handleConfirmation(msg)
		}
		return m.handleSelection(msg)
	}

	return m, nil
}

// handleSelection handles repository selection keys
func (m *RepositorySelectorModel) handleSelection(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		m.done = true
		return m, tea.Quit

	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < len(m.repositories)-1 {
			m.cursor++
		}

	case " ":
		// Toggle selection
		if m.selected[m.cursor] {
			delete(m.selected, m.cursor)
		} else {
			m.selected[m.cursor] = true
		}

	case "a":
		// Select all
		for i := range m.repositories {
			m.selected[i] = true
		}

	case "n":
		// Select none
		m.selected = make(map[int]bool)

	case "enter":
		if len(m.selected) > 0 {
			m.showConfirm = true
		}
	}

	return m, nil
}

// handleConfirmation handles the confirmation dialog
func (m *RepositorySelectorModel) handleConfirmation(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "y", "Y":
		m.confirmed = true
		m.done = true
		return m, tea.Quit

	case "n", "N", "esc":
		m.showConfirm = false

	case "q", "ctrl+c":
		m.done = true
		return m, tea.Quit
	}

	return m, nil
}

// View renders the UI
func (m *RepositorySelectorModel) View() string {
	if m.showConfirm {
		return m.renderConfirmation()
	}

	return m.renderSelection()
}

// renderSelection renders the repository selection view
func (m *RepositorySelectorModel) renderSelection() string {
	var s strings.Builder

	// Title
	title := titleStyle.Render("🚀 Mass Git Cloner - Repository Selection")
	s.WriteString(title + "\n\n")

	// Header with filter info and selection count
	headerText := fmt.Sprintf("%s - %d selected", m.filter.String(), len(m.selected))
	header := headerStyle.Render(headerText)
	s.WriteString(header + "\n\n")

	// Calculate visible range for scrolling
	visibleHeight := m.height - 8 // Reserve space for header, footer, etc.
	start := 0
	end := len(m.repositories)

	if len(m.repositories) > visibleHeight {
		start = m.cursor - visibleHeight/2
		if start < 0 {
			start = 0
		}
		end = start + visibleHeight
		if end > len(m.repositories) {
			end = len(m.repositories)
			start = end - visibleHeight
			if start < 0 {
				start = 0
			}
		}
	}

	// Repository list
	for i := start; i < end; i++ {
		repo := m.repositories[i]
		
		// Cursor indicator
		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render("❯")
		}

		// Selection checkbox
		checkbox := "☐"
		checkStyle := uncheckedStyle
		if m.selected[i] {
			checkbox = "☑"
			checkStyle = checkedStyle
		}

		// Repository name
		repoName := repo.Name
		if len(repoName) > 25 {
			repoName = repoName[:22] + "..."
		}

		// Language tag
		language := repo.Language
		if language == "" {
			language = "N/A"
		}
		langTag := languageStyle.Render(language)

		// Stars and forks
		stars := starsStyle.Render(fmt.Sprintf("★ %d", repo.StarCount))
		forks := forksStyle.Render(fmt.Sprintf("⑂ %d", repo.ForkCount))

		// Description
		description := repo.Description
		if len(description) > 40 {
			description = description[:37] + "..."
		}
		if description == "" {
			description = "No description"
		}

		// Build the line
		line := fmt.Sprintf("%s %s %-25s %s %s %s %s",
			cursor,
			checkStyle.Render(checkbox),
			repoName,
			langTag,
			stars,
			forks,
			description,
		)

		// Highlight the current line
		if m.cursor == i {
			line = selectedStyle.Render(line)
		}

		s.WriteString(line + "\n")
	}

	// Show scrolling indicator if needed
	if len(m.repositories) > visibleHeight {
		scrollInfo := fmt.Sprintf("\n📄 Showing %d-%d of %d repositories", start+1, end, len(m.repositories))
		s.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280")).Render(scrollInfo) + "\n")
	}

	// Help text
	help := helpStyle.Render(`
Controls:
  ↑/k: Move up    ↓/j: Move down    Space: Toggle selection
  a: Select all   n: Select none    Enter: Confirm selection
  q: Quit`)

	s.WriteString(help)

	return s.String()
}

// renderConfirmation renders the confirmation dialog
func (m *RepositorySelectorModel) renderConfirmation() string {
	var s strings.Builder

	title := titleStyle.Render("🚀 Confirm Repository Selection")
	s.WriteString(title + "\n\n")

	// Show selected repositories
	s.WriteString(fmt.Sprintf("You have selected %s repositories for cloning:\n\n", 
		confirmStyle.Render(fmt.Sprintf("%d", len(m.selected)))))

	// List selected repositories
	count := 0
	for i, repo := range m.repositories {
		if m.selected[i] {
			count++
			langTag := ""
			if repo.Language != "" {
				langTag = languageStyle.Render(repo.Language)
			}
			
			line := fmt.Sprintf("  %d. %s %s %s %s",
				count,
				repo.Name,
				langTag,
				starsStyle.Render(fmt.Sprintf("★ %d", repo.StarCount)),
				forksStyle.Render(fmt.Sprintf("⑂ %d", repo.ForkCount)),
			)
			s.WriteString(line + "\n")
		}
	}

	s.WriteString("\n")
	
	// Confirmation prompt
	confirmPrompt := confirmStyle.Render("Do you want to proceed with cloning these repositories? (y/N)")
	s.WriteString(confirmPrompt + "\n\n")

	help := helpStyle.Render("y: Yes, clone them    n/Esc: Go back    q: Quit")
	s.WriteString(help)

	return s.String()
}

// GetSelectedRepositories returns the selected repositories
func (m *RepositorySelectorModel) GetSelectedRepositories() []*models.Repository {
	var selected []*models.Repository
	for i := range m.selected {
		if i < len(m.repositories) {
			selected = append(selected, m.repositories[i])
		}
	}
	return selected
}

// IsConfirmed returns whether the user confirmed the selection
func (m *RepositorySelectorModel) IsConfirmed() bool {
	return m.confirmed
}

// IsDone returns whether the selection is complete
func (m *RepositorySelectorModel) IsDone() bool {
	return m.done
}

// ShowRepositorySelector shows the interactive repository selector
func ShowRepositorySelector(repos []*models.Repository, filter models.FilterType) ([]*models.Repository, error) {
	if len(repos) == 0 {
		return nil, fmt.Errorf("no repositories to select from")
	}

	model := NewRepositorySelectorModel(repos, filter)
	
	program := tea.NewProgram(model, tea.WithAltScreen())
	
	finalModel, err := program.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run repository selector: %w", err)
	}

	selectorModel := finalModel.(*RepositorySelectorModel)
	
	if !selectorModel.IsConfirmed() {
		return nil, fmt.Errorf("selection cancelled by user")
	}

	return selectorModel.GetSelectedRepositories(), nil
}