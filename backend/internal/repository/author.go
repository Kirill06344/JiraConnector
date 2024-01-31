package repository

import (
	"backend/internal/entity"
	"gorm.io/gorm"
)

type AuthorRepositoryImpl struct {
	db *gorm.DB
}

func newAuthorRepository(db *gorm.DB) *AuthorRepositoryImpl {
	return &AuthorRepositoryImpl{db: db}
}

func (repo *AuthorRepositoryImpl) FindById(id uint) (*entity.Author, error) {
	var author entity.Author
	result := repo.db.First(&author, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &author, nil
}
