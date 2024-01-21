package service

import (
	"backend/internal/dto"
	"backend/internal/repository"
)

type ProjectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) *ProjectService {
	return &ProjectService{
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

func (s *ProjectService) Create(project *dto.Project) error {
	return s.repo.Create(project)
}

func (s *ProjectService) Update(id uint, project *dto.Project) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	project.Id = id
	return s.repo.Update(project)
}

func (s *ProjectService) Delete(id uint) error {
	return s.repo.Delete(id)
}
