package main

import (
	"log"

	"github.com/Desgue/Tasker-Cli/tui/form"
	"github.com/Desgue/Tasker-Cli/tui/message"
	"github.com/Desgue/Tasker-Cli/tui/project"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	projects State = iota
	projectForm
)

type State int

type model struct {
	currentState State
	models       []tea.Model
}

func New(state State) *model {
	model := &model{}
	model.currentState = state
	model.models = []tea.Model{project.New(), form.NewProjectForm()}
	return model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	case message.ProjectForm:
		m.currentState = projectForm
		return m, nil
	}
	switch m.currentState {
	case projects:
		m.models[m.currentState], cmd = m.models[m.currentState].Update(msg)
		return m, cmd
	case projectForm:
		m.models[m.currentState], cmd = m.models[m.currentState].Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	return m.models[m.currentState].View()
}

func main() {
	m := New(projects)

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
