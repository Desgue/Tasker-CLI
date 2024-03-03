package repo

import (
	"fmt"

	"github.com/Desgue/Tasker-Cli/domain"
	"github.com/Desgue/Tasker-Cli/repo/db"
)

type TaskRepository interface {
	CreateTask(domain.TaskRequest) (domain.TaskResponse, error)
	GetTasks(int) ([]domain.TaskResponse, error)
	DeleteTask(int) error
	UpdateTask(domain.TaskRequest) (domain.TaskResponse, error)
}

type taskRepository struct {
	sql *db.SqliteDB
}

func NewTaskRepository(db *db.SqliteDB) TaskRepository {
	return &taskRepository{sql: db}
}

func (r *taskRepository) CreateTask(t domain.TaskRequest) (domain.TaskResponse, error) {
	result, err := r.sql.DB.Exec("INSERT INTO Tasks (projectId, title, description, status) VALUES (?, ?, ?, ?)", t.ProjectId, t.Title, t.Description, t.Status)
	if err != nil {
		return domain.TaskResponse{}, err
	}
	createdId, err := result.LastInsertId()
	if err != nil {
		return domain.TaskResponse{}, err
	}
	Id := int(createdId)

	var response domain.TaskResponse

	err = r.sql.DB.QueryRow("SELECT * FROM Tasks WHERE id = ?", Id).Scan(&response.Id, &response.ProjectId, &response.Title, &response.Description, &response.Status)
	if err != nil {
		return domain.TaskResponse{}, err
	}

	return response, nil
}

func (r *taskRepository) GetTasks(projectId int) ([]domain.TaskResponse, error) {
	rows, err := r.sql.DB.Query("SELECT * FROM Tasks where projectId = ?", projectId)
	if err != nil {
		return nil, fmt.Errorf("error at repository fetching tasks: %s", err)
	}
	defer rows.Close()
	var tasks []domain.TaskResponse
	for rows.Next() {
		var t domain.TaskResponse
		err := rows.Scan(&t.Id, &t.ProjectId, &t.Title, &t.Description, &t.Status)
		if err != nil {
			return nil, fmt.Errorf("error at repository scanning tasks: %s", err)
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *taskRepository) DeleteTask(id int) error {
	_, err := r.sql.DB.Exec("DELETE FROM Tasks WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) UpdateTask(t domain.TaskRequest) (domain.TaskResponse, error) {
	_, err := r.sql.DB.Exec("UPDATE Tasks SET title = ?, description = ?, status = ? WHERE id = ?", t.Title, t.Description, t.Status, t.Id)
	if err != nil {
		return domain.TaskResponse{}, err
	}
	var response domain.TaskResponse
	err = r.sql.DB.QueryRow("SELECT * FROM Tasks WHERE id = ?", t.Id).Scan(&response.Id, &response.Title, &response.Description, &response.Status)
	if err != nil {
		return domain.TaskResponse{}, err
	}
	return response, nil
}
