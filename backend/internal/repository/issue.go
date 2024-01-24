package repository

import (
	"backend/internal/entity"
	"gorm.io/gorm"
)

type IssueRepositoryImpl struct {
	db *gorm.DB
}

func newIssueRepository(db *gorm.DB) *IssueRepositoryImpl {
	return &IssueRepositoryImpl{db: db}
}

func (repo *IssueRepositoryImpl) Find() ([]entity.Issue, error) {
	var issues []entity.Issue
	result := repo.db.Find(&issues)
	if result.Error != nil {
		return nil, result.Error
	}
	return issues, nil
}

func (repo *IssueRepositoryImpl) FindById(id uint) (*entity.Issue, error) {
	var issue entity.Issue
	result := repo.db.First(&issue, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &issue, nil
}

func (repo *IssueRepositoryImpl) Create(issue *entity.Issue) error {
	result := repo.db.Create(issue)
	return result.Error
}

func (repo *IssueRepositoryImpl) Update(issue *entity.Issue) error {
	result := repo.db.Save(&issue)
	return result.Error
}

func (repo *IssueRepositoryImpl) Delete(id uint) error {
	result := repo.db.Delete(&entity.Issue{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
