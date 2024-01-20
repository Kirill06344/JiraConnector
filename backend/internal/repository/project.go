package repository

import (
	"backend/internal/dto"
	"backend/internal/entity"
	"backend/internal/service"
	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) service.ProjectRepository {
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

func (repo *ProjectRepositoryImpl) Create(project *dto.Project) error {
	//TODO implement me
	panic("implement me")
}

func (repo *ProjectRepositoryImpl) Update(id uint, issue *dto.Project) error {
	//TODO implement me
	panic("implement me")
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
