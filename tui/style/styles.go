package style

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	BorderColor lipgloss.Color
	Focused     lipgloss.Style
	Column      lipgloss.Style
	Help        lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)
	s.Focused = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#6200EE"))
	s.Column = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("0"))
	s.Help = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	return s
}

type FormStyle struct {
	TextInput lipgloss.Style
	TextArea  lipgloss.Style
}

func DefaultFormStyle() *FormStyle {
	s := new(FormStyle)
	s.TextInput = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#6200EE"))
	s.TextInput.Width(50)
	s.TextArea = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#6200EE"))
	s.TextArea.Width(50)
	return s
}
