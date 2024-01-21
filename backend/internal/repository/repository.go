package repository

import "gorm.io/gorm"

type Repository struct {
	Issue   IssueRepository
	Project ProjectRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Issue:   newIssueRepository(db),
		Project: newProjectRepository(db),
	}
}
