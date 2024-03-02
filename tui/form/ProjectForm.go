package form

import (
	"log"

	"github.com/Desgue/Tasker-Cli/domain"
	"github.com/Desgue/Tasker-Cli/repo"
	"github.com/Desgue/Tasker-Cli/repo/db"
	"github.com/Desgue/Tasker-Cli/svc"
	"github.com/Desgue/Tasker-Cli/tui/message"
	"github.com/Desgue/Tasker-Cli/tui/style"
	"github.com/Desgue/Tasker-Cli/types"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ProjectForm is a form for creating a new project
type ProjectForm struct {
	service     svc.ProjectService
	title       textinput.Model
	description textarea.Model
	styles      *style.FormStyle
	focused     types.Priority
	Width       int
	Height      int
}

func NewProjectForm(db *db.SqliteDB) *ProjectForm {
	repository := repo.NewProjectRepository(db)
	service := svc.NewProjectService(repository)
	form := &ProjectForm{styles: style.DefaultFormStyle(), service: service}
	form.defaultConfig()
	return form

}

func (m ProjectForm) Init() tea.Cmd {
	return nil
}

func (m ProjectForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case message.ShowProjectForm:
		m.focused = msg.Focused
		m.Width = msg.Width
		m.Height = msg.Height
		log.Println(m.focused)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "crtk+r":
			m.reset()
		case "esc":
			return m, m.GoToProjectList
		case "tab":
			if m.title.Focused() {
				m.next()
				return m, textarea.Blink
			}
			m.previous()
			return m, textinput.Blink
		case "enter":
			if m.title.Focused() {
				m.next()
				return m, textinput.Blink
			} else {
				return m, m.CreateProject
			}
		}
	}
	var cmds []tea.Cmd
	m.title, cmd = m.title.Update(msg)
	cmds = append(cmds, cmd)
	m.description, cmd = m.description.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m ProjectForm) View() string {
	m.styles.TextInput.Width(50)
	m.styles.TextInput.Height(1)

	m.styles.TextArea.Width(m.Width / 4)
	m.styles.TextArea.Height(m.Height / 4)
	titleView := m.styles.TextInput.Render(m.title.View())
	descriptionView := m.styles.TextArea.Render(m.description.View())

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, lipgloss.JoinVertical(
		lipgloss.Center,
		titleView,
		descriptionView,
	))
}

// HELPERS
func (form *ProjectForm) defaultConfig() {
	form.title = textinput.New()
	form.title.Placeholder = "Add a title for the project"
	form.title.Cursor.Blink = true
	form.title.Focus()
	form.description = textarea.New()
	form.description.Placeholder = "Add a description for the project"
}

func (m *ProjectForm) next() {
	m.title.Blur()
	m.description.Focus()
}
func (m *ProjectForm) previous() {
	m.description.Blur()
	m.title.Focus()
}
func (m *ProjectForm) reset() {
	m.title.Focus()
	m.title.SetValue("")
	m.description.Blur()
	m.description.SetValue("")

}
func (m ProjectForm) GoToProjectList() tea.Msg {
	return message.ShowProjectList{}

}
func (m ProjectForm) CreateProject() tea.Msg {
	p := domain.NewProjectRequest(m.title.Value(), m.description.Value(), m.focused)
	createdProject, err := m.service.AddProject(p)
	if err != nil {
		log.Println(err)
		return err
	}
	return createdProject

}
