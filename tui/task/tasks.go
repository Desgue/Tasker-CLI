package task

import (
	"log"

	"github.com/Desgue/Tasker-Cli/domain"
	"github.com/Desgue/Tasker-Cli/repo"
	"github.com/Desgue/Tasker-Cli/repo/db"
	"github.com/Desgue/Tasker-Cli/svc"
	"github.com/Desgue/Tasker-Cli/tui/message"
	"github.com/Desgue/Tasker-Cli/tui/style"
	"github.com/Desgue/Tasker-Cli/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor int = 4
const (
	Pending types.Status = iota
	InProgress
	Done
)

type Model struct {
	service   svc.TaskService
	ProjectId int
	Lists     []list.Model
	Focused   types.Status
	styles    *style.Styles
	width     int
	height    int
}

func New(db *db.SqliteDB) *Model {
	repo := repo.NewTaskRepository(db)
	service := svc.NewTaskService(repo)
	m := &Model{styles: style.DefaultStyles(), service: service, Focused: Pending}
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case message.ShowTaskList:
		m.width = msg.Width
		m.height = msg.Height
		m.styles.Column.Width(msg.Width / divisor)
		m.styles.Focused.Width(msg.Width / divisor)
		m.ProjectId = msg.ProjectId
		m.initLists(msg.Width, msg.Height)
	case domain.TaskItem:
		m.initLists(m.width, m.height)
		m, cmd := m.Update(nil)
		return m, cmd
	case tea.KeyMsg:
		switch msg.String() {
		// QUIT PROGRAM
		case "q", "ctrl+c":
			return m, tea.Quit
		// GO TO PROJECTS KANBANVIEW
		case "esc":
			return m, m.GoToProjects
		// NAVIGATE BETWEEN BOARDS
		case "l", "right", "tab":
			m.nextBoardView()
		case "h", "left":
			m.previousBoardView()

		// GO TO TASK FORM
		case "n":
			return m, m.GoToForm

		// UPDATE TASK STATUS
		case "ctrl+n", " ":
			if m.Lists[m.Focused].SelectedItem() != nil {
				m.moveStatusForward()
			}
			m.initLists(m.width, m.height)
			return m.Update(nil)
		case "backspace", "ctrl+b":
			if m.Lists[m.Focused].SelectedItem() != nil {
				m.moveStatusBackward()
			}
			m.initLists(m.width, m.height)
			return m.Update(nil)
		}
	}
	var cmd tea.Cmd
	m.Lists[m.Focused], cmd = m.Lists[m.Focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	todoView := m.Lists[Pending].View()
	inProgressView := m.Lists[InProgress].View()
	doneView := m.Lists[Done].View()
	switch m.Focused {
	default:
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.styles.Focused.Render(todoView),
			m.styles.Column.Render(inProgressView),
			m.styles.Column.Render(doneView),
		))
	case InProgress:
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.styles.Column.Render(todoView),
			m.styles.Focused.Render(inProgressView),
			m.styles.Column.Render(doneView),
		))
	case Done:
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.styles.Column.Render(todoView),
			m.styles.Column.Render(inProgressView),
			m.styles.Focused.Render(doneView),
		))
	}
}

// HELPERS
func (m *Model) initLists(w, h int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), w, h-divisor)
	defaultList.SetStatusBarItemName("Task", "Tasks")
	defaultList.SetShowHelp(false)
	m.Lists = []list.Model{defaultList, defaultList, defaultList}
	for i := 0; i < len(m.Lists); i++ {
		m.Lists[i].Title = types.Status(i).String()
	}
	m.fetchTasks()
}
func (m *Model) fetchTasks() {
	tasks, err := m.service.GetTasks(m.ProjectId)
	if err != nil {
		log.Fatalln("error fetching tasks: ", err)
	}
	for _, t := range tasks {
		m.Lists[t.Status].InsertItem(0, t)
	}

}

func (m *Model) moveStatusForward() {
	selectedTask := m.Lists[m.Focused].SelectedItem()
	task := selectedTask.(domain.TaskItem)
	task.Next()
	req := domain.TaskRequestFromItem(task)
	_, err := m.service.UpdateTask(req)
	if err != nil {
		log.Println("error updating task: ", err)
	}
}

func (m *Model) moveStatusBackward() {
	selectedTask := m.Lists[m.Focused].SelectedItem()
	task := selectedTask.(domain.TaskItem)
	task.Previous()
	req := domain.TaskRequestFromItem(task)
	_, err := m.service.UpdateTask(req)
	if err != nil {
		log.Println("error updating task: ", err)
	}
}

func (m Model) GoToForm() tea.Msg {
	return message.ShowTaskForm{
		Focused:   m.Focused,
		ProjectId: m.ProjectId,
		Width:     m.width,
		Height:    m.height}
}

func (m Model) GoToProjects() tea.Msg {
	return message.ShowProjectList{}
}

func (m *Model) nextBoardView() {
	if m.Focused < Done {
		m.Focused++
	} else {
		m.Focused = Pending
	}

}

func (m *Model) previousBoardView() {
	if m.Focused > Pending {
		m.Focused--
	} else {
		m.Focused = Done
	}

}
