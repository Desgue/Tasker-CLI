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
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			return m, m.GoToProjects
		case "l", "right", "tab":
			m.next()
		case "h", "left":
			m.previous()
		case "n":
			return m, m.GoToForm

		}
		return m, nil

	}
	return m, nil
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

func (m *Model) next() {
	if m.Focused < Done {
		m.Focused++
	} else {
		m.Focused = Pending
	}

}

func (m *Model) previous() {
	if m.Focused > Pending {
		m.Focused--
	} else {
		m.Focused = Done
	}

}
