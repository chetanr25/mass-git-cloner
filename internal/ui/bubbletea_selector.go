package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7C3AED")).
			Padding(0, 1).
			Bold(true)

	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#3B82F6")).
			Padding(0, 1)

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

	languageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#8B5CF6")).
			Padding(0, 1).
			MarginRight(1)

	starsStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FBBF24"))

	forksStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#34D399"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280")).
			Margin(1, 0)

	confirmStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#EF4444")).
			Padding(0, 1).
			Bold(true)
)

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

func (m *RepositorySelectorModel) Init() tea.Cmd {
	return nil
}

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
		if m.selected[m.cursor] {
			delete(m.selected, m.cursor)
		} else {
			m.selected[m.cursor] = true
		}

	case "a":
		for i := range m.repositories {
			m.selected[i] = true
		}

	case "n":
		m.selected = make(map[int]bool)

	case "enter":
		if len(m.selected) > 0 {
			m.showConfirm = true
		}
	}
	return m, nil
}

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

func (m *RepositorySelectorModel) View() string {
	if m.showConfirm {
		return m.renderConfirmation()
	}

	return m.renderSelection()
}

func (m *RepositorySelectorModel) renderSelection() string {
	var s strings.Builder

	title := titleStyle.Render("🚀 Mass Git Cloner - Repository Selection")
	s.WriteString(title + "\n\n")

	headerText := fmt.Sprintf("%s - %d selected", m.filter.String(), len(m.selected))
	header := headerStyle.Render(headerText)
	s.WriteString(header + "\n\n")

	visibleHeight := m.height - 8
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

	for i := start; i < end; i++ {
		repo := m.repositories[i]

		cursor := " "
		if m.cursor == i {
			cursor = cursorStyle.Render("❯")
		}

		checkbox := "☐"
		checkStyle := uncheckedStyle
		if m.selected[i] {
			checkbox = "✓"
			checkStyle = checkedStyle
		}

		repoName := repo.Name
		if len(repoName) > 25 {
			repoName = repoName[:22] + "..."
		}

		language := repo.Language
		if language == "" {
			language = "N/A"
		}
		langTag := languageStyle.Render(language)

		stars := starsStyle.Render(fmt.Sprintf("★ %d", repo.StarCount))
		forks := forksStyle.Render(fmt.Sprintf("⑂ %d", repo.ForkCount))

		description := repo.Description
		if len(description) > 40 {
			description = description[:37] + "..."
		}
		if description == "" {
			description = "No description"
		}

		line := fmt.Sprintf("%s %s %-25s %s %s %s %s",
			cursor,
			checkStyle.Render(checkbox),
			repoName,
			langTag,
			stars,
			forks,
			description,
		)

		if m.cursor == i {
			line = selectedStyle.Render(line)
		}

		s.WriteString(line + "\n")
	}

	if len(m.repositories) > visibleHeight {
		scrollInfo := fmt.Sprintf("\n📄 Showing %d-%d of %d repositories", start+1, end, len(m.repositories))
		s.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280")).Render(scrollInfo) + "\n")
	}

	help := helpStyle.Render(`
Controls:
  ↑/k: Move up    ↓/j: Move down    Space: Toggle selection
  a: Select all   n: Select none    Enter: Confirm selection
  q: Quit`)

	s.WriteString(help)

	return s.String()
}

func (m *RepositorySelectorModel) renderConfirmation() string {
	var s strings.Builder

	title := titleStyle.Render("🚀 Confirm Repository Selection")
	s.WriteString(title + "\n\n")

	s.WriteString(fmt.Sprintf("You have selected %s repositories for cloning:\n\n",
		confirmStyle.Render(fmt.Sprintf("%d", len(m.selected)))))

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

	confirmPrompt := confirmStyle.Render("Do you want to proceed with cloning these repositories? (y/N)")
	s.WriteString(confirmPrompt + "\n\n")

	help := helpStyle.Render("y: Yes, clone them    n/Esc: Go back    q: Quit")
	s.WriteString(help)

	return s.String()
}

func (m *RepositorySelectorModel) GetSelectedRepositories() []*models.Repository {
	var selected []*models.Repository
	for i := range m.selected {
		if i < len(m.repositories) {
			selected = append(selected, m.repositories[i])
		}
	}
	return selected
}

func (m *RepositorySelectorModel) IsConfirmed() bool {
	return m.confirmed
}

func (m *RepositorySelectorModel) IsDone() bool {
	return m.done
}

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
