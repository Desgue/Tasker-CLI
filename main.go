package main

import (
	"log"

	"github.com/Desgue/Tasker-Cli/cfg"
	"github.com/Desgue/Tasker-Cli/repo"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	//m := New(projects)

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("Error creating log file: %v", err)
	}
	defer f.Close()

	dir := cfg.SetupPath()
	_, err = repo.Open(dir)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	/* p := tea.NewProgram(m, tea.WithAltScreen())
	_, err = p.Run()
	if err != nil {
		log.Fatalf("Error running program: %v", err)
	} */

}
