package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

type StatsDisplayModel struct {
	stats    *models.RepositoryStats
	username string
	done     bool
}

func NewStatsDisplayModel(stats *models.RepositoryStats, username string) *StatsDisplayModel {
	return &StatsDisplayModel{
		stats:    stats,
		username: username,
		done:     false,
	}
}

func (m *StatsDisplayModel) Init() tea.Cmd {
	return nil
}

func (m *StatsDisplayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.done = true
			return m, tea.Quit

		case "enter", " ":
			m.done = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *StatsDisplayModel) View() string {
	var s strings.Builder

	title := titleStyle.Render(fmt.Sprintf("🚀 Mass Git Cloner - %s", m.username))
	s.WriteString(title + "\n\n")

	successMsg := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#10B981")).
		Padding(0, 1).
		Bold(true).
		Render(fmt.Sprintf("✅ Found %d repositories", m.stats.Total))
	s.WriteString(successMsg + "\n\n")

	statsContainer := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#3B82F6")).
		Padding(1, 2).
		Margin(1, 0)

	totalStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#3B82F6")).
		Padding(0, 1).
		Bold(true)

	nonForkStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#10B981")).
		Padding(0, 1).
		Bold(true)

	forkStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#F59E0B")).
		Padding(0, 1).
		Bold(true)

	publicStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#8B5CF6")).
		Padding(0, 1).
		Bold(true)

	privateStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#EF4444")).
		Padding(0, 1).
		Bold(true)

	var statsContent strings.Builder
	statsContent.WriteString("📊 Repository Statistics\n\n")

	totalBar := strings.Repeat("█", min(m.stats.Total, 50))
	nonForkBar := strings.Repeat("█", min(m.stats.NonForks, 50))
	forkBar := strings.Repeat("█", min(m.stats.Forks, 50))

	statsContent.WriteString(fmt.Sprintf("Total repositories:      %s %d\n",
		totalStyle.Render(totalBar), m.stats.Total))
	statsContent.WriteString(fmt.Sprintf("Original repositories:   %s %d\n",
		nonForkStyle.Render(nonForkBar), m.stats.NonForks))
	statsContent.WriteString(fmt.Sprintf("Forked repositories:     %s %d\n",
		forkStyle.Render(forkBar), m.stats.Forks))
	statsContent.WriteString(fmt.Sprintf("Public repositories:     %s %d\n",
		publicStyle.Render("████████████████████"), m.stats.Public))
	statsContent.WriteString(fmt.Sprintf("Private repositories:    %s %d\n",
		privateStyle.Render("████████████████████"), m.stats.Private))

	if m.stats.Total > 0 {
		nonForkPct := float64(m.stats.NonForks) / float64(m.stats.Total) * 100
		forkPct := float64(m.stats.Forks) / float64(m.stats.Total) * 100

		statsContent.WriteString("\n Breakdown:\n")
		statsContent.WriteString(fmt.Sprintf("   Original: %.1f%%\n", nonForkPct))
		statsContent.WriteString(fmt.Sprintf("   Forks: %.1f%%\n", forkPct))
	}

	s.WriteString(statsContainer.Render(statsContent.String()) + "\n")

	help := helpStyle.Render("Press Enter or Space to continue, q to quit")
	s.WriteString(help)

	return s.String()
}

func (m *StatsDisplayModel) IsDone() bool {
	return m.done
}

func ShowRepositoryStats(stats *models.RepositoryStats, username string) error {
	model := NewStatsDisplayModel(stats, username)

	program := tea.NewProgram(model, tea.WithAltScreen())

	_, err := program.Run()
	if err != nil {
		return fmt.Errorf("failed to display statistics: %w", err)
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
