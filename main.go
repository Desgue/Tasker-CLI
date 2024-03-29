package main

import (
	"log"

	"github.com/Desgue/Tasker-Cli/cfg"
	"github.com/Desgue/Tasker-Cli/repo/db"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	dir := cfg.SetupPath()
	db, err := db.Open(dir)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	m := New(projects, db)

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
