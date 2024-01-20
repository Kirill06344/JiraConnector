package repository

import (
	"backend/internal/dto"
	"backend/internal/entity"
	"gorm.io/gorm"
)

type IssueRepositoryImpl struct {
	db *gorm.DB
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

func (repo *IssueRepositoryImpl) Create(id uint, issue *dto.Issue) error {
	//TODO implement me
	return nil
}

func (repo *IssueRepositoryImpl) Update(issue *dto.Issue) error {
	//TODO implement me
	return nil
}

func (repo *IssueRepositoryImpl) Delete(id uint) error {
	result := repo.db.Delete(&entity.Issue{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
