package repository

import (
	"backend/internal/entity"
	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct {
	db *gorm.DB
}

func newProjectRepository(db *gorm.DB) *ProjectRepositoryImpl {
	return &ProjectRepositoryImpl{db: db}
}

func (repo *ProjectRepositoryImpl) Find() ([]entity.Project, error) {
	var projects []entity.Project
	result := repo.db.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (repo *ProjectRepositoryImpl) FindById(id uint) (*entity.Project, error) {
	var project entity.Project
	result := repo.db.First(&project, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func (repo *ProjectRepositoryImpl) Create(project *entity.Project) error {
	result := repo.db.Create(&project)
	return result.Error
}

func (repo *ProjectRepositoryImpl) Update(project *entity.Project) error {
	result := repo.db.Save(&project)
	return result.Error
}

func (repo *ProjectRepositoryImpl) Delete(id uint) error {
	result := repo.db.Delete(&entity.Project{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
