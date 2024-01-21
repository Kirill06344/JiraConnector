package service

import "backend/internal/repository"

type Service struct {
	Project *ProjectService
	Issue   *IssueService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Project: NewProjectService(repository.Project),
		Issue:   NewIssueService(repository.Issue),
	}
}
