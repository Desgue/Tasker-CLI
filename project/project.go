package project

import "time"

type Project struct {
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
