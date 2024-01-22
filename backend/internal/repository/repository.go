package repository

import "gorm.io/gorm"

type Repository struct {
	Issue   IssueRepository
	Project ProjectRepository
	Author  AuthorRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Issue:   newIssueRepository(db),
		Project: newProjectRepository(db),
		Author:  newAuthorRepository(db),
	}
}
