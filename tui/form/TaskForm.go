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

type TaskForm struct {
	ProjectId   int
	service     svc.TaskService
	title       textinput.Model
	description textarea.Model
	styles      *style.FormStyle
	status      types.Status
	Width       int
	Height      int
}

func NewTaskForm(db *db.SqliteDB) *TaskForm {
	repository := repo.NewTaskRepository(db)
	service := svc.NewTaskService(repository)
	form := &TaskForm{styles: style.DefaultFormStyle(), service: service}
	form.defaultConfig()
	return form
}

// Implement the Model interface
func (form TaskForm) Init() tea.Cmd {
	return nil
}

func (form TaskForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case message.ShowTaskForm:
		form.ProjectId = msg.ProjectId
		form.status = msg.Focused
		form.Width = msg.Width
		form.Height = msg.Height
		return form, textinput.Blink
	case tea.KeyMsg:
		switch msg.String() {
		//Quit
		case "ctrl+c":
			return form, tea.Quit
			// Reset fields
		case "crtk+r":
			form.reset()
			// Go to KanbanBoard
		case "esc":
			return form, form.GoToTaskList
			// Move between fields
		case "tab":
			if form.title.Focused() {
				form.next()
				return form, textarea.Blink
			}
			form.previous()
			return form, textinput.Blink
			// Move from title to description or save if description is focused
		case "ctrl+s":
			if form.title.Focused() {
				form.next()
				return form, textarea.Blink
			}
			return form, form.CreateTask

		}
	}
	var cmds []tea.Cmd
	form.title, cmd = form.title.Update(msg)
	cmds = append(cmds, cmd)
	form.description, cmd = form.description.Update(msg)
	cmds = append(cmds, cmd)
	return form, tea.Batch(cmds...)
}

func (form TaskForm) View() string {
	titleView := form.styles.TextInput.Render(form.title.View())
	descriptionView := form.styles.TextArea.Render(form.description.View())
	return lipgloss.Place(form.Width, form.Height, lipgloss.Center, lipgloss.Center, lipgloss.JoinVertical(
		lipgloss.Center,
		titleView,
		descriptionView,
	))
}

// HELPERS
func (form *TaskForm) defaultConfig() {
	form.title = textinput.New()
	form.title.Placeholder = "Add a title for the task"
	form.title.Cursor.Blink = true
	form.title.Focus()
	form.description = textarea.New()
	form.description.Placeholder = "Add a description"
}

func (form *TaskForm) reset() {
	form.title.Reset()
	form.title.Focus()
	form.description.Reset()
}
func (form TaskForm) GoToTaskList() tea.Msg {
	return message.ShowTaskList{ProjectId: form.ProjectId, Width: form.Width, Height: form.Height}
}

func (form *TaskForm) next() {
	form.title.Blur()
	form.description.Focus()
}

func (form *TaskForm) previous() {
	form.description.Blur()
	form.title.Focus()
}

func (form *TaskForm) CreateTask() tea.Msg {
	newTask := domain.NewTaskRequest(form.ProjectId, form.title.Value(), form.description.Value(), form.status)
	res, err := form.service.AddTask(newTask)
	if err != nil {
		log.Println(err)
		return err
	}
	return res

}
