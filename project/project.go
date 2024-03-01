package project

import "time"

const (
	Low Priority = iota
	Medium
	High
)

type Priority int

func (p Priority) String() string {
	return [3]string{"Low", "Medium", "High"}[p]
}

type Project struct {
	title       string
	description string
	Priority    Priority
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProject(title, description string, Priority Priority) *Project {
	return &Project{
		title:       title,
		description: description,
		Priority:    Priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

//Implement list.Item interface
func (p Project) Title() string {
	return p.title
}
func (p Project) Description() string {
	return p.description
}
func (p Project) FilterValue() string {
	return p.title
}

//HELPERS
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
