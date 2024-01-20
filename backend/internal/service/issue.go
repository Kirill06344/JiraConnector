package service

import (
	"backend/internal/dto"
	"backend/internal/entity"
)

type IssueRepository interface {
	Find() ([]entity.Issue, error)
	FindById(id uint) (entity.Issue, error)
	Create(issue *dto.Issue) error
	Update(id uint, issue *dto.Issue) error
	Delete(id uint) error
}

type IssueService struct {
	repo IssueRepository
}

func NewIssueService(repo IssueRepository) IssueService {
	return IssueService{
		repo: repo,
	}
}
