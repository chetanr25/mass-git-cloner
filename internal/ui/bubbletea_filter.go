package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

type FilterOption struct {
	Filter      models.FilterType
	Name        string
	Description string
	Count       int
}

type FilterSelectorModel struct {
	options  []FilterOption
	cursor   int
	selected models.FilterType
	done     bool
}

func NewFilterSelectorModel(stats *models.RepositoryStats) *FilterSelectorModel {
	options := []FilterOption{
		{
			Filter:      models.FilterAll,
			Name:        "All Repositories",
			Description: "Include both original and forked repositories",
			Count:       stats.Total,
		},
		{
			Filter:      models.FilterNonForks,
			Name:        "Original Repositories Only",
			Description: "Exclude forked repositories",
			Count:       stats.NonForks,
		},
		{
			Filter:      models.FilterForksOnly,
			Name:        "Forked Repositories Only",
			Description: "Include only forked repositories",
			Count:       stats.Forks,
		},
	}

	return &FilterSelectorModel{
		options:  options,
		cursor:   0,
		selected: models.FilterAll,
		done:     false,
	}
}

func (m *FilterSelectorModel) Init() tea.Cmd {
	return nil
}

func (m *FilterSelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.done = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.selected = m.options[m.cursor].Filter
			m.done = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *FilterSelectorModel) View() string {
	var s strings.Builder

	title := titleStyle.Render("🚀 Mass Git Cloner - Filter Selection")
	s.WriteString(title + "\n\n")

	header := headerStyle.Render("Select the type of repositories to clone:")
	s.WriteString(header + "\n\n")

	for i, option := range m.options {
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("❯ ")
		}

		optionStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA"))
		countStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#3B82F6")).
			Padding(0, 1).
			Bold(true)
		descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#9CA3AF"))

		if m.cursor == i {
			optionStyle = selectedStyle
		}

		count := countStyle.Render(fmt.Sprintf("%d", option.Count))
		name := optionStyle.Render(option.Name)
		desc := descStyle.Render(option.Description)

		line := fmt.Sprintf("%s%s %s\n    %s", cursor, name, count, desc)
		s.WriteString(line + "\n\n")
	}

	help := helpStyle.Render(`
Controls:
  ↑/k: Move up    ↓/j: Move down    Enter/Space: Select    q: Quit`)

	s.WriteString(help)

	return s.String()
}

func (m *FilterSelectorModel) GetSelectedFilter() models.FilterType {
	return m.selected
}

func (m *FilterSelectorModel) IsDone() bool {
	return m.done
}

func ShowFilterSelector(stats *models.RepositoryStats) (models.FilterType, error) {
	model := NewFilterSelectorModel(stats)

	program := tea.NewProgram(model, tea.WithAltScreen())

	finalModel, err := program.Run()
	if err != nil {
		return models.FilterAll, fmt.Errorf("failed to run filter selector: %w", err)
	}

	filterModel := finalModel.(*FilterSelectorModel)

	if !filterModel.IsDone() {
		return models.FilterAll, fmt.Errorf("selection cancelled by user")
	}

	return filterModel.GetSelectedFilter(), nil
}
