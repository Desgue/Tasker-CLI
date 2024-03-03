package main

import (
	"github.com/Desgue/Tasker-Cli/domain"
	"github.com/Desgue/Tasker-Cli/repo/db"
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
	taskForm
)

type State int

type model struct {
	db           *db.SqliteDB
	currentState State
	models       map[State]tea.Model
}

func New(state State, db *db.SqliteDB) *model {
	model := &model{db: db}
	model.currentState = state
	model.models = map[State]tea.Model{
		projects:    project.New(db),
		projectForm: form.NewProjectForm(db),
		tasks:       task.New(db),
		taskForm:    form.NewTaskForm(db),
	}

	return model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	// STATE CHANGE MESSAGES
	case message.ShowProjectForm:
		m.currentState = projectForm
		m.models[projectForm], cmd = form.NewProjectForm(m.db).Update(msg)
		return m, cmd
	case message.ShowProjectList:
		m.currentState = projects
		return m, nil
	case domain.ProjectItem:
		m.currentState = projects
		m.models[projects], cmd = m.models[projects].Update(msg)
		return m, cmd
	case message.ShowTaskList:
		m.currentState = tasks
		m.models[tasks], cmd = m.models[tasks].Update(msg)
		return m, cmd
	case message.ShowTaskForm:
		m.currentState = taskForm
		m.models[taskForm], cmd = form.NewTaskForm(m.db).Update(msg)
		return m, cmd
	case domain.TaskItem:
		m.currentState = tasks
		m.models[tasks], cmd = m.models[tasks].Update(msg)
		return m, cmd

	}
	//
	switch m.currentState {
	case projects:
		m.models[m.currentState], cmd = m.models[m.currentState].Update(msg)
		return m, cmd
	case projectForm:
		m.models[m.currentState], cmd = m.models[m.currentState].Update(msg)
		return m, cmd
	case tasks:
		m.models[m.currentState], cmd = m.models[m.currentState].Update(msg)
		return m, cmd
	case taskForm:
		m.models[m.currentState], cmd = m.models[m.currentState].Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	return m.models[m.currentState].View()
}
