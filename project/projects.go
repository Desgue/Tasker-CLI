package project

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const divisor int = 4

type Model struct {
	Lists   []list.Model
	Err     error
	Focused Priority
	loaded  bool
}

func New() *Model {
	m := &Model{}
	return m
}

func (m *Model) InitLists(w, h int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), w, h)
	m.Lists = []list.Model{defaultList, defaultList, defaultList}

	// Init Low Priority
	m.Lists[Low].Title = "Low Priority"
	m.Lists[Low].SetItems([]list.Item{
		NewProject("Project 1", "Low Priority Project 1", Low),
		NewProject("Project 2", "Low Priority Project 2", Low),
		NewProject("Project 3", "Low Priority Project 3", Low),
	})
	m.Lists[Low].SetStatusBarItemName("Project", "Projects")

	// Init Medium Priority
	m.Lists[Medium].Title = "Medium Priority"
	m.Lists[Medium].SetItems([]list.Item{
		NewProject("Project 4", "Medium Priority Project 1", Medium),
		NewProject("Project 5", "Medium Priority Project 2", Medium),
		NewProject("Project 6", "Medium Priority Project 3", Medium),
	})
	m.Lists[Medium].SetStatusBarItemName("Project", "Projects")

	// Init High Priority
	m.Lists[High].Title = "High Priority"
	m.Lists[High].SetItems([]list.Item{
		NewProject("Project 7", "High Priority Project 1", High),
		NewProject("Project 8", "High Priority Project 2", High),
		NewProject("Project 9", "High Priority Project 3", High),
	})
	m.Lists[High].SetStatusBarItemName("Project", "Projects")

}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.InitLists(msg.Width/divisor, msg.Height)
		m.loaded = true
	case tea.KeyMsg:
		switch msg.String() {
		case "l", "right":
			if m.Focused < High {
				m.Focused++
			} else if m.Focused == High {
				m.Focused = Low
			}
		case "h", "left":
			if m.Focused > Low {
				m.Focused--
			} else if m.Focused == Low {
				m.Focused = High
			}
		}
	}

	var cmd tea.Cmd
	m.Lists[m.Focused], cmd = m.Lists[m.Focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if !m.loaded {
		return "Loading..."
	}
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.Lists[Low].View(),
		m.Lists[Medium].View(),
		m.Lists[High].View(),
	)
}
