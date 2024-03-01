package project

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
	Err     error
	Focused types.Priority
	loaded  bool
	styles  *style.Styles
	width   int
	height  int
}

func New() *Model {
	m := &Model{styles: style.DefaultStyles()}
	return m
}

func (m *Model) InitLists(w, h int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), w, h-divisor)
	defaultList.SetStatusBarItemName("Project", "Projects")
	defaultList.SetShowHelp(false)
	m.Lists = []list.Model{defaultList, defaultList, defaultList}

	// Init Low Priority
	m.Lists[Low].Title = "Low Priority"
	m.Lists[Low].SetItems([]list.Item{
		NewProject("Project 1", "Low Priority Project 1 ", Low),
		NewProject("Project 2", "Low Priority Project 2", Low),
		NewProject("Project 3", "Low Priority Project 3", Low),
	})

	// Init Medium Priority
	m.Lists[Medium].Title = "Medium Priority"
	m.Lists[Medium].SetItems([]list.Item{
		NewProject("Project 4", "Medium Priority Project 1", Medium),
		NewProject("Project 5", "Medium Priority Project 2", Medium),
		NewProject("Project 6", "Medium Priority Project 3", Medium),
	})

	// Init High Priority
	m.Lists[High].Title = "High Priority"
	m.Lists[High].SetItems([]list.Item{
		NewProject("Project 7", "High Priority Project 1", High),
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
	case Project:
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
		}
	}

	var cmd tea.Cmd
	m.Lists[m.Focused], cmd = m.Lists[m.Focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.loaded {
		lowView := m.Lists[Low].View()
		mediumView := m.Lists[Medium].View()
		highView := m.Lists[High].View()
		switch m.Focused {
		default:
			return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, lipgloss.JoinHorizontal(
				lipgloss.Left,
				m.styles.Focused.Render(lowView),
				m.styles.Column.Render(mediumView),
				m.styles.Column.Render(highView),
			))
		case Medium:
			return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, lipgloss.JoinHorizontal(
				lipgloss.Left,
				m.styles.Column.Render(lowView),
				m.styles.Focused.Render(mediumView),
				m.styles.Column.Render(highView),
			))
		case High:
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
func (m Model) CreateProject(p Project) (Project, error) {
	return p, nil
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
	selectedProject := selectedItem.(Project)
	m.Lists[selectedProject.Priority].RemoveItem(m.Lists[m.Focused].Index())
	selectedProject.next()
	m.Lists[selectedProject.Priority].InsertItem(0, selectedProject)

}

func (m *Model) moveToPrevious() /* tea.Msg */ {
	selectedItem := m.Lists[m.Focused].SelectedItem()
	selectedProject := selectedItem.(Project)
	m.Lists[selectedProject.Priority].RemoveItem(m.Lists[m.Focused].Index())
	selectedProject.previous()
	m.Lists[selectedProject.Priority].InsertItem(0, selectedProject)

}

func (m *Model) next() {
	if m.Focused < High {
		m.Focused++
	} else {
		m.Focused = Low
	}
}
func (m *Model) previous() {
	if m.Focused > Low {
		m.Focused--
	} else {
		m.Focused = High
	}
}
