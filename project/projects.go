package project

import (
	"github.com/Desgue/Tasker-Cli/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Projects []Project
	Index    int
	Width    int
	Height   int
	Styles   *style.Styles
}

func New(p []Project) *Model {
	styles := style.DefaultStyles()
	return &Model{
		Projects: p,
		Styles:   styles,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.Width == 0 {
		return "LOADING..."
	}
	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, "Project View")
}
