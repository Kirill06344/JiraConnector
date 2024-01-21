package repository

import (
	"backend/internal/dto"
	"backend/internal/entity"
)

type IssueRepository interface {
	Find() ([]entity.Issue, error)
	FindById(id uint) (*entity.Issue, error)
	Create(issue *dto.Issue) error
	Update(issue *dto.Issue) error
	Delete(id uint) error
}

type ProjectRepository interface {
	Find() ([]entity.Project, error)
	FindById(id uint) (*entity.Project, error)
	Create(project *dto.Project) error
	Update(project *dto.Project) error
	Delete(id uint) error
}
