package svc

import (
	"github.com/Desgue/Tasker-Cli/domain"
	"github.com/Desgue/Tasker-Cli/repo"
)

type TaskService interface {
	AddProject(domain.TaskRequest) (domain.TaskItem, error)
	GetTasks() ([]domain.TaskItem, error)
	DeleteTask(int) error
	UpdateTask(domain.TaskRequest) (domain.TaskItem, error)
}

type taskService struct {
	repo repo.TaskRepository
}

func NewTaskService(repo repo.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) AddProject(t domain.TaskRequest) (domain.TaskItem, error) {
	taskRes, err := s.repo.CreateTask(t)
	if err != nil {
		return domain.TaskItem{}, err
	}
	taskItem := domain.NewTaskItem(taskRes)
	return taskItem, nil
}

func (s *taskService) GetTasks() ([]domain.TaskItem, error) {
	res, err := s.repo.GetTasks()
	if err != nil {
		return nil, err
	}
	var tasks []domain.TaskItem
	for _, t := range res {
		tasks = append(tasks, domain.NewTaskItem(t))
	}
	return tasks, nil
}

func (s *taskService) DeleteTask(id int) error {
	err := s.repo.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *taskService) UpdateTask(t domain.TaskRequest) (domain.TaskItem, error) {
	taskRes, err := s.repo.UpdateTask(t)
	if err != nil {
		return domain.TaskItem{}, err
	}
	taskItem := domain.NewTaskItem(taskRes)
	return taskItem, nil
}
