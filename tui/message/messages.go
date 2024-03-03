package message

import (
	"github.com/Desgue/Tasker-Cli/domain"
	"github.com/Desgue/Tasker-Cli/types"
)

type ShowPreviousPage struct{}
type ShowProjectList struct{}
type ShowProjectForm struct {
	Focused types.Priority
	Width   int
	Height  int
}

type ShowTaskList struct {
	Focused   types.Status
	ProjectId int
	Width     int
	Height    int
}
type ShowTaskForm struct {
	Task      domain.TaskItem
	ProjectId int
	Focused   types.Status
	Width     int
	Height    int
}
