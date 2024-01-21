package service

import "backend/internal/repository"

type IssueService struct {
	repo repository.IssueRepository
}

func NewIssueService(repo repository.IssueRepository) *IssueService {
	return &IssueService{
		repo: repo,
	}
}
