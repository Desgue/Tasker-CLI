package task

import (
	"time"

	"github.com/Desgue/Tasker-Cli/types"
)

const (
	Todo types.Status = iota
	InProgress
	Done
)

type Task struct {
	title       string
	description string
	Status      types.Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(title, description string) *Task {
	return &Task{
		title:       title,
		description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// Implement the list.Item interface

func (t Task) Title() string {
	return t.title
}
func (t Task) Description() string {
	return t.description
}
func (t Task) FilterValue() string {
	return t.description
}
