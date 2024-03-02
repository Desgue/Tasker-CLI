package domain

import (
	"time"

	"github.com/Desgue/Tasker-Cli/types"
)

type ProjectRequest struct {
	Id          int            `json:"id" db:"id"`
	Title       string         `json:"title" db:"title"`
	Description string         `json:"description" db:"description"`
	Priority    types.Priority `json:"priority" db:"priority"`
	CreatedAt   time.Time      `json:"createdAt" db:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt" db:"updatedAt"`
}

const (
	Low types.Priority = iota
	Medium
	High
)

type ProjectItem struct {
	Id          int
	title       string
	description string
	Priority    types.Priority
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProject(title, description string, Priority types.Priority) ProjectItem {
	return ProjectItem{
		title:       title,
		description: description,
		Priority:    Priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// Implement list.Item interface
func (p ProjectItem) Title() string {
	return p.title
}
func (p ProjectItem) Description() string {
	return p.description
}
func (p ProjectItem) FilterValue() string {
	return p.title
}

// HELPERS
func (p *ProjectItem) Next() {
	if p.Priority < High {
		p.Priority++
	} else {
		p.Priority = Low
	}
}
func (p *ProjectItem) Previous() {
	if p.Priority > Low {
		p.Priority--
	} else {
		p.Priority = High
	}
}
