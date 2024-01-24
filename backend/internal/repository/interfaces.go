package repository

import (
	"backend/internal/entity"
)

type IssueRepository interface {
	Find() ([]entity.Issue, error)
	FindById(id uint) (*entity.Issue, error)
	Create(issue *entity.Issue) error
	Update(issue *entity.Issue) error
	Delete(id uint) error
}

type ProjectRepository interface {
	Find() ([]entity.Project, error)
	FindById(id uint) (*entity.Project, error)
	Create(project *entity.Project) error
	Update(project *entity.Project) error
	Delete(id uint) error
}

type AuthorRepository interface {
	FindById(id uint) (*entity.Author, error)
}
