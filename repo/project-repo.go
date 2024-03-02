package repo

import "github.com/Desgue/Tasker-Cli/repo/db"

type ProjectRepository interface {
	CreateProject()
}

type projectRepository struct {
	repo *db.SqliteDB
}
