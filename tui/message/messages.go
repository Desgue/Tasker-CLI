package message

import tea "github.com/charmbracelet/bubbletea"

type ProjectForm struct{}

func ShowProjectForm() tea.Msg {
	return ProjectForm{}
}
