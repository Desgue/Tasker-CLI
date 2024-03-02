package project

import (
	"github.com/Desgue/Tasker-Cli/domain"
	"github.com/Desgue/Tasker-Cli/repo/db"
	"github.com/Desgue/Tasker-Cli/tui/message"
	"github.com/Desgue/Tasker-Cli/tui/style"
	"github.com/Desgue/Tasker-Cli/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor int = 4

type Model struct {
	repo    *db.SqliteDB
	Lists   []list.Model
	Err     error
	Focused types.Priority
	loaded  bool
	styles  *style.Styles
	width   int
	height  int
}

func New(repo *db.SqliteDB) *Model {
	m := &Model{styles: style.DefaultStyles(), repo: repo, Focused: domain.Low}
	return m
}

func (m *Model) InitLists(w, h int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), w, h-divisor)
	defaultList.SetStatusBarItemName("Project", "Projects")
	defaultList.SetShowHelp(false)
	m.Lists = []list.Model{defaultList, defaultList, defaultList}

	// Init domain.Low Priority
	m.Lists[domain.Low].Title = "domain.Low Priority"
	m.Lists[domain.Low].SetItems([]list.Item{
		domain.NewProject("Project 1", "domain.Low Priority Project 1 ", domain.Low),
		domain.NewProject("Project 2", "domain.Low Priority Project 2", domain.Low),
		domain.NewProject("Project 3", "domain.Low Priority Project 3", domain.Low),
	})

	// Init domain.Medium Priority
	m.Lists[domain.Medium].Title = "domain.Medium Priority"
	m.Lists[domain.Medium].SetItems([]list.Item{
		domain.NewProject("Project 4", "domain.Medium Priority Project 1", domain.Medium),
		domain.NewProject("Project 5", "domain.Medium Priority Project 2", domain.Medium),
		domain.NewProject("Project 6", "domain.Medium Priority Project 3", domain.Medium),
	})

	// Init domain.High Priority
	m.Lists[domain.High].Title = "domain.High Priority"
	m.Lists[domain.High].SetItems([]list.Item{
		domain.NewProject("Project 7", "domain.High Priority Project 1", domain.High),
	})

}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.styles.Column.Width(msg.Width / divisor)
		m.styles.Focused.Width(msg.Width / divisor)
		m.width = msg.Width
		m.height = msg.Height
		m.InitLists(msg.Width, msg.Height)
		m.loaded = true
	case domain.ProjectItem:
		m.Lists[msg.Priority].InsertItem(0, msg)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "l", "right":
			m.next()
		case "h", "left":
			m.previous()
		case "space", "enter":
			m.moveToNext()
		case "backspace":
			m.moveToPrevious()
		case "n":
			return m, m.GoToForm
		case "t":
			return m, m.GoToTasks
		}
	}

	var cmd tea.Cmd
	m.Lists[m.Focused], cmd = m.Lists[m.Focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.loaded {
		lowView := m.Lists[domain.Low].View()
		mediumView := m.Lists[domain.Medium].View()
		highView := m.Lists[domain.High].View()
		switch m.Focused {
		default:
			return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, lipgloss.JoinHorizontal(
				lipgloss.Left,
				m.styles.Focused.Render(lowView),
				m.styles.Column.Render(mediumView),
				m.styles.Column.Render(highView),
			))
		case domain.Medium:
			return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, lipgloss.JoinHorizontal(
				lipgloss.Left,
				m.styles.Column.Render(lowView),
				m.styles.Focused.Render(mediumView),
				m.styles.Column.Render(highView),
			))
		case domain.High:
			return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, lipgloss.JoinHorizontal(
				lipgloss.Left,
				m.styles.Column.Render(lowView),
				m.styles.Column.Render(mediumView),
				m.styles.Focused.Render(highView),
			))
		}

	}

	return "Loading..."
}

// HELPERS
func (m Model) CreateProject(p domain.ProjectItem) (domain.ProjectItem, error) {
	return p, nil
}

func (m Model) GoToTasks() tea.Msg {
	id := m.Lists[m.Focused].SelectedItem().(domain.ProjectItem).Id
	return message.ShowTaskList{
		ProjectId: id,
		Width:     m.width,
		Height:    m.height,
	}
}

func (m Model) GoToForm() tea.Msg {
	msg := message.ShowProjectForm{
		Focused: m.Focused,
		Width:   m.width,
		Height:  m.height,
	}
	return msg
}
func (m *Model) moveToNext() /* tea.Msg */ {
	selectedItem := m.Lists[m.Focused].SelectedItem()
	selectedProject := selectedItem.(domain.ProjectItem)
	m.Lists[selectedProject.Priority].RemoveItem(m.Lists[m.Focused].Index())
	selectedProject.Next()
	m.Lists[selectedProject.Priority].InsertItem(0, selectedProject)

}

func (m *Model) moveToPrevious() /* tea.Msg */ {
	selectedItem := m.Lists[m.Focused].SelectedItem()
	selectedProject := selectedItem.(domain.ProjectItem)
	m.Lists[selectedProject.Priority].RemoveItem(m.Lists[m.Focused].Index())
	selectedProject.Previous()
	m.Lists[selectedProject.Priority].InsertItem(0, selectedProject)

}

func (m *Model) next() {
	if m.Focused < domain.High {
		m.Focused++
	} else {
		m.Focused = domain.Low
	}
}
func (m *Model) previous() {
	if m.Focused > domain.Low {
		m.Focused--
	} else {
		m.Focused = domain.High
	}
}
