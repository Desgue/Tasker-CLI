package svc

import "github.com/Desgue/Tasker-Cli/repo"

type ProjectService interface {
	AddProject()
}

type projectService struct {
	repo repo.ProjectRepository
}

func NewProjectService(repo repo.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) AddProject() {
	s.repo.CreateProject()
}
