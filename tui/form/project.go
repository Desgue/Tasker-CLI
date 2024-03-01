package form

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Project is a form for creating a new project
type Project struct {
	title       textinput.Model
	description textarea.Model
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
	return m, nil
}

func (m Project) View() string {
	return "Project form"
}
