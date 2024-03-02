package main

import (
	"github.com/Desgue/Tasker-Cli/tui/form"
	"github.com/Desgue/Tasker-Cli/tui/message"
	"github.com/Desgue/Tasker-Cli/tui/project"
	"github.com/Desgue/Tasker-Cli/tui/task"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	projects State = iota
	projectForm
	tasks
)

type State int

type model struct {
	currentState State
	models       []tea.Model
}

func New(state State) *model {
	model := &model{}
	model.currentState = state
	model.models = []tea.Model{project.New(), form.NewProjectForm(), task.New()}
	return model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.(type) {
	// STATE CHANGE MESSAGES
	case message.ShowProjectForm:
		m.currentState = projectForm
		m.models[projectForm], cmd = form.NewProjectForm().Update(msg)
		return m, cmd
	case message.ShowProjectList:
		m.currentState = projects
		return m, nil
	case message.ShowTaskList:
		m.currentState = tasks
		m.models[tasks], cmd = m.models[tasks].Update(msg)
		return m, cmd
	case project.Project:
		m.currentState = projects
		m.models[projects], cmd = m.models[projects].Update(msg)
		return m, cmd

	}
	//
	switch m.currentState {
	case projects:
		m.models[projects], cmd = m.models[projects].Update(msg)
		return m, cmd
	case projectForm:
		m.models[projectForm], cmd = m.models[projectForm].Update(msg)
		return m, cmd
	case tasks:
		m.models[tasks], cmd = m.models[tasks].Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	return m.models[m.currentState].View()
}
