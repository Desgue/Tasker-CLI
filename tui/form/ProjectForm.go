package form

import (
	"github.com/Desgue/Tasker-Cli/tui/message"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Project is a form for creating a new project
type Project struct {
	title       textinput.Model
	description textarea.Model
	width       int
	height      int
}

func NewProjectForm() *Project {
	form := &Project{}
	form.title = textinput.New()
	form.title.Placeholder = "Add a title for the project"
	form.title.Focus()
	form.description = textarea.New()
	form.description.Placeholder = "Add a description for the project"
	return form

}

func (m Project) Init() tea.Cmd {
	return nil
}

func (m Project) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "backspace":
			return m, m.GoToProjectList
		case "enter":
			if m.title.Focused() {
				m.title.Blur()
				m.description.Focus()
				return m, textinput.Blink
			} else {
				return m, m.NewProject
			}
		}
	}
	return m, nil
}

func (m Project) View() string {
	if m.width == 0 {
		return "Loading..."

	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.title.View(),
			m.description.View(),
		))
}

// HELPERS

func (m Project) GoToProjectList() tea.Msg {
	return message.ShowPreviousPage{}

}
func (m Project) NewProject() tea.Msg {
	return message.ShowProjectList{}

}
