package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/chetanr25/mass-git-cloner/pkg/models"
)

// CloneProgressModel represents the Bubbletea model for cloning progress
type CloneProgressModel struct {
	progress    *models.CloneProgress
	startTime   time.Time
	completed   []*models.CloneResult
	failed      []*models.CloneResult
	currentRepo string
	done        bool
}

// NewCloneProgressModel creates a new clone progress model
func NewCloneProgressModel(total int) *CloneProgressModel {
	return &CloneProgressModel{
		progress: &models.CloneProgress{
			Total:     total,
			Completed: 0,
			Failed:    0,
			Current:   "",
		},
		startTime:   time.Now(),
		completed:   make([]*models.CloneResult, 0),
		failed:      make([]*models.CloneResult, 0),
		currentRepo: "",
		done:        false,
	}
}

// Init initializes the model
func (m *CloneProgressModel) Init() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return TickMsg{Time: t}
	})
}

// TickMsg represents a tick message for updating the progress
type TickMsg struct {
	Time time.Time
}

// CloneUpdateMsg represents a clone update message
type CloneUpdateMsg struct {
	Result *models.CloneResult
}

// CurrentRepoMsg represents the current repository being cloned
type CurrentRepoMsg struct {
	RepoName string
}

// DoneMsg represents completion of all cloning
type DoneMsg struct{}

// Update handles messages and updates the model
func (m *CloneProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			if m.done {
				return m, tea.Quit
			}
		}

	case TickMsg:
		if !m.done {
			return m, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
				return TickMsg{Time: t}
			})
		}

	case CurrentRepoMsg:
		m.currentRepo = msg.RepoName
		m.progress.Current = msg.RepoName

	case CloneUpdateMsg:
		if msg.Result.Success {
			m.completed = append(m.completed, msg.Result)
			m.progress.Completed++
		} else {
			m.failed = append(m.failed, msg.Result)
			m.progress.Failed++
		}

	case DoneMsg:
		m.done = true
	}

	return m, nil
}

// View renders the cloning progress
func (m *CloneProgressModel) View() string {
	var s strings.Builder

	// Title
	title := titleStyle.Render("🚀 Mass Git Cloner - Cloning Progress")
	s.WriteString(title + "\n\n")

	// Progress bar
	total := m.progress.Total
	completed := m.progress.Completed + m.progress.Failed
	percentage := 0.0
	if total > 0 {
		percentage = float64(completed) / float64(total) * 100
	}

	// Create visual progress bar
	barWidth := 50
	filled := int(float64(barWidth) * percentage / 100)
	progressBar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

	// Progress bar styling
	progressStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#10B981")).
		Bold(true)

	progressInfo := fmt.Sprintf("[%s] %.1f%% (%d/%d)",
		progressStyle.Render(progressBar), percentage, completed, total)

	s.WriteString(progressInfo + "\n\n")

	// Stats
	elapsed := time.Since(m.startTime).Truncate(time.Second)
	statsContainer := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#3B82F6")).
		Padding(0, 2)

	var statsContent strings.Builder
	statsContent.WriteString(fmt.Sprintf("⏱️  Elapsed: %s\n", elapsed))
	statsContent.WriteString(fmt.Sprintf("✅ Completed: %s\n", 
		checkedStyle.Render(fmt.Sprintf("%d", m.progress.Completed))))
	statsContent.WriteString(fmt.Sprintf("❌ Failed: %s\n", 
		lipgloss.NewStyle().Foreground(lipgloss.Color("#EF4444")).Render(fmt.Sprintf("%d", m.progress.Failed))))

	if m.progress.Completed > 0 && total > 0 {
		successRate := float64(m.progress.Completed) / float64(completed) * 100
		statsContent.WriteString(fmt.Sprintf("📈 Success Rate: %.1f%%", successRate))
	}

	s.WriteString(statsContainer.Render(statsContent.String()) + "\n\n")

	// Current operation
	if m.currentRepo != "" && !m.done {
		currentStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#F59E0B")).
			Padding(0, 1).
			Bold(true)
		
		current := currentStyle.Render(fmt.Sprintf("⏳ Cloning: %s", m.currentRepo))
		s.WriteString(current + "\n\n")
	}

	// Recent completions (last 3)
	if len(m.completed) > 0 {
		s.WriteString("📋 Recent completions:\n")
		start := len(m.completed) - 3
		if start < 0 {
			start = 0
		}
		
		for i := start; i < len(m.completed); i++ {
			result := m.completed[i]
			duration := result.Duration.Truncate(time.Millisecond * 100)
			line := fmt.Sprintf("  ✅ %s (%s)", result.Repository.Name, duration)
			s.WriteString(checkedStyle.Render(line) + "\n")
		}
		s.WriteString("\n")
	}

	// Recent failures (last 2)
	if len(m.failed) > 0 {
		s.WriteString("📋 Recent failures:\n")
		start := len(m.failed) - 2
		if start < 0 {
			start = 0
		}
		
		failStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#EF4444"))
		for i := start; i < len(m.failed); i++ {
			result := m.failed[i]
			line := fmt.Sprintf("  ❌ %s: %s", result.Repository.Name, result.Error.Error())
			if len(line) > 80 {
				line = line[:77] + "..."
			}
			s.WriteString(failStyle.Render(line) + "\n")
		}
		s.WriteString("\n")
	}

	// Completion message or help
	if m.done {
		completionStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#10B981")).
			Padding(0, 1).
			Bold(true)
		
		completion := completionStyle.Render("🎉 Cloning completed! Press q to exit.")
		s.WriteString(completion + "\n")
	} else {
		help := helpStyle.Render("Press Ctrl+C to cancel")
		s.WriteString(help)
	}

	return s.String()
}

// IsDone returns whether the progress is complete
func (m *CloneProgressModel) IsDone() bool {
	return m.done
}

// GetResults returns the cloning results
func (m *CloneProgressModel) GetResults() ([]*models.CloneResult, []*models.CloneResult) {
	return m.completed, m.failed
}

// UpdateProgress updates the progress with a new result
func (m *CloneProgressModel) UpdateProgress(result *models.CloneResult) tea.Cmd {
	return func() tea.Msg {
		return CloneUpdateMsg{Result: result}
	}
}

// SetCurrentRepo sets the current repository being cloned
func (m *CloneProgressModel) SetCurrentRepo(repoName string) tea.Cmd {
	return func() tea.Msg {
		return CurrentRepoMsg{RepoName: repoName}
	}
}

// MarkDone marks the progress as complete
func (m *CloneProgressModel) MarkDone() tea.Cmd {
	return func() tea.Msg {
		return DoneMsg{}
	}
}