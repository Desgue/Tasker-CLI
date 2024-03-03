package domain

import (
	"log"

	"github.com/Desgue/Tasker-Cli/types"
)

const (
	Pending types.Status = iota
	InProgress
	Done
)

type TaskRequest struct {
	Id          int    `json:"id" db:"id"`
	ProjectId   int    `json:"projectId" db:"projectId"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Status      string `json:"status" db:"status"`
}

type TaskResponse struct {
	Id          int    `json:"id" db:"id"`
	ProjectId   int    `json:"projectId" db:"projectId"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Status      string `json:"status" db:"status"`
}

func TaskResponseFromItem(t TaskItem) TaskResponse {
	return TaskResponse{
		Id:          t.Id,
		ProjectId:   t.ProjectId,
		Title:       t.title,
		Description: t.description,
		Status:      t.Status.String(),
	}
}

func NewTaskRequest(projectId int, title, description string, Status types.Status) TaskRequest {
	StatusStr := Status.String()
	return TaskRequest{
		ProjectId:   projectId,
		Title:       title,
		Description: description,
		Status:      StatusStr,
	}
}

func TaskRequestFromItem(t TaskItem) TaskRequest {
	return TaskRequest{
		Id:          t.Id,
		ProjectId:   t.ProjectId,
		Title:       t.title,
		Description: t.description,
		Status:      t.Status.String(),
	}
}

type TaskItem struct {
	Id          int
	ProjectId   int
	title       string
	description string
	Status      types.Status
}

func NewTaskItem(t TaskResponse) TaskItem {
	status, err := types.StrToStatus(t.Status)
	if err != nil {
		log.Printf("Error: %s, formating status to default value of Pending", err)
		status = Pending
	}
	return TaskItem{
		Id:          t.Id,
		title:       t.Title,
		description: t.Description,
		Status:      status,
	}
}

// Implement list.Item interface
func (t TaskItem) Title() string {
	return t.title
}

func (t TaskItem) Description() string {
	return t.description
}

func (t TaskItem) FilterValue() string {
	return t.title
}

// Method to change the status of a task

func (t *TaskItem) Next() {
	if t.Status < Done {
		t.Status++
	} else {
		t.Status = Pending
	}
}

func (t *TaskItem) Previous() {
	if t.Status > Pending {
		t.Status--
	} else {
		t.Status = Done
	}
}
