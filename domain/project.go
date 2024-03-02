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
