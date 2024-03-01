package message

import (
	"github.com/Desgue/Tasker-Cli/types"
)

type ShowPreviousPage struct{}
type ShowProjectList struct{}
type ShowProjectForm struct {
	Focused types.Priority
	Width   int
	Height  int
}
