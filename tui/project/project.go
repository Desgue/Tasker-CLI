package project

import (
	"time"

	"github.com/Desgue/Tasker-Cli/types"
)

const (
	Low types.Priority = iota
	Medium
	High
)

type Project struct {
	title       string
	description string
	Priority    types.Priority
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProject(title, description string, Priority types.Priority) Project {
	return Project{
		title:       title,
		description: description,
		Priority:    Priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// Implement list.Item interface
func (p Project) Title() string {
	return p.title
}
func (p Project) Description() string {
	return p.description
}
func (p Project) FilterValue() string {
	return p.title
}

// HELPERS
func (p *Project) next() {
	if p.Priority < High {
		p.Priority++
	} else {
		p.Priority = Low
	}
}
func (p *Project) previous() {
	if p.Priority > Low {
		p.Priority--
	} else {
		p.Priority = High
	}
}
