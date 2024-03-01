package task

import (
	"github.com/Desgue/Tasker-Cli/tui/message"
	"github.com/Desgue/Tasker-Cli/tui/style"
	"github.com/Desgue/Tasker-Cli/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor int = 4

type Model struct {
	Lists   []list.Model
	Focused types.Status
	styles  *style.Styles
	width   int
	height  int
}

func New() *Model {
	m := &Model{styles: style.DefaultStyles()}
	return m
}

func (m *Model) InitLists(w, h, projectId int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), w, h-divisor)
	defaultList.SetStatusBarItemName("Task", "Tasks")
	defaultList.SetShowHelp(false)
	m.Lists = []list.Model{defaultList, defaultList, defaultList}
	// TODO: Implement function that retrieves tasks from database

	// Init Todo
	m.Lists[Todo].Title = Todo.String()

	// Init InProgress
	m.Lists[InProgress].Title = InProgress.String()

	// Init Done
	m.Lists[Done].Title = Done.String()
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
		projectId := msg.ProjectId
		m.InitLists(msg.Width, msg.Height, projectId)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			return m, m.GoToProjects
		case "l", "right":
			m.next()
		case "h", "left":
			m.previous()

		}
		return m, nil

	}
	return m, nil
}

func (m Model) View() string {
	todoView := m.Lists[Todo].View()
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

func (m Model) GoToProjects() tea.Msg {
	return message.ShowProjectList{}
}

func (m *Model) next() {
	if m.Focused < Done {
		m.Focused++
	} else {
		m.Focused = Todo
	}

}

func (m *Model) previous() {
	if m.Focused > Todo {
		m.Focused--
	} else {
		m.Focused = Done
	}

}
