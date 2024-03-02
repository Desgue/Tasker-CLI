package domain

import (
	"github.com/Desgue/Tasker-Cli/types"
)

const (
	Low types.Priority = iota
	Medium
	High
)

type ProjectRequest struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Priority    string `json:"priority" db:"priority"`
}

func NewProjectRequest(title, description string, Priority types.Priority) ProjectRequest {
	PriorityStr := Priority.String()
	return ProjectRequest{
		Title:       title,
		Description: description,
		Priority:    PriorityStr,
	}
}

type ProjectItem struct {
	Id          int
	title       string
	description string
	Priority    types.Priority
}

func NewProjectItem(p ProjectRequest) ProjectItem {
	Priority, err := types.StrToPriority(p.Priority)
	if err != nil {
		Priority = Low
	}
	return ProjectItem{
		title:       p.Title,
		description: p.Description,
		Priority:    Priority,
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
