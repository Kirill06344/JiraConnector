package service

import (
	"backend/internal/dto"
	"backend/internal/entity"
)

type ProjectRepository interface {
	Find() ([]entity.Project, error)
	FindById(id uint) (*entity.Project, error)
	Create(issue *dto.Project) error
	Update(id uint, issue *dto.Project) error
	Delete(id uint) error
}

type ProjectService struct {
	repo ProjectRepository
}

func NewProjectService(repo ProjectRepository) ProjectService {
	return ProjectService{
		repo: repo,
	}
}

func (s *ProjectService) Find() ([]dto.Project, error) {
	data, err := s.repo.Find()
	if err != nil {
		return nil, err
	}
	projects := make([]dto.Project, len(data))
	for i, el := range data {
		projects[i] = dto.Project{
			Id:    el.Id,
			Title: el.Title,
		}
	}
	return projects, err
}

func (s *ProjectService) FindById(id uint) (*dto.Project, error) {
	data, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return &dto.Project{
		Id:    data.Id,
		Title: data.Title,
	}, nil
}

func (s *ProjectService) Delete(id uint) error {
	return s.repo.Delete(id)
}
