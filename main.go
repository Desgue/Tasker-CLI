package main

import (
	"log"

	"github.com/Desgue/Tasker-Cli/project"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	projectView State = iota
)

type State int

type model struct {
	currentState State
	projectView  *project.Model
}

func New(state State) *model {

	return &model{
		currentState: state,
		projectView:  project.New([]project.Project{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.currentState {
	case projectView:
		return m.projectView.Update(msg)
	}
	return m, nil
}

func (m model) View() string {
	switch m.currentState {
	case projectView:
		return m.projectView.View()
	}
	return "Main model view"
}

func main() {
	m := New(projectView)

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("Error creating log file: %v", err)
	}
	defer f.Close()

	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		log.Fatalf("Error running program: %v", err)
	}

}
