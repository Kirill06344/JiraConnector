package service

import (
	"backend/internal/dto"
	"backend/internal/entity"
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
			Id:          el.ID,
			Description: el.Title,
			Key:         el.Key,
			Name:        el.Name,
		}
	}
	return projects, nil
}

func (s *ProjectService) FindById(id uint) (*dto.Project, error) {
	data, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return &dto.Project{
		Id:          data.ID,
		Description: data.Title,
		Key:         data.Key,
		Name:        data.Name,
	}, nil
}

func (s *ProjectService) Create(project *dto.Project) error {
	var model = entity.Project{Title: project.Description, Key: project.Key}
	return s.repo.Create(&model)
}

func (s *ProjectService) Update(id uint, project *dto.Project) error {
	_, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	project.Id = id
	var model = entity.Project{ID: project.Id, Title: project.Description, Key: project.Key}
	return s.repo.Update(&model)
}

func (s *ProjectService) Delete(id uint) error {
	return s.repo.Delete(id)
}
