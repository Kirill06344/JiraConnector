package service

import (
	"backend/internal/dto"
	"backend/internal/entity"
	"backend/internal/repository"
)

type IssueService struct {
	repo *repository.Repository
}

func NewIssueService(repo *repository.Repository) *IssueService {
	return &IssueService{
		repo: repo,
	}
}

func (s *IssueService) Find() ([]dto.Issue, error) {
	data, err := s.repo.Issue.Find()
	if err != nil {
		return nil, err
	}
	issues := make([]dto.Issue, len(data))
	for i, el := range data {
		issues[i] = s.convertEntityToDto(&el)
	}
	return issues, nil
}

func (s *IssueService) FindById(id uint) (*dto.Issue, error) {
	data, err := s.repo.Issue.FindById(id)
	if err != nil {
		return nil, err
	}
	result := s.convertEntityToDto(data)
	return &result, nil
}

func (s *IssueService) Create(issue *dto.Issue) error {
	//TODO
	panic("implement me")
}

func (s *IssueService) Update(id uint, issue *dto.Issue) error {
	//TODO
	panic("implement me")
}

func (s *IssueService) Delete(id uint) error {
	return s.repo.Issue.Delete(id)
}

func (s *IssueService) convertEntityToDto(el *entity.Issue) dto.Issue {
	project, _ := s.repo.Project.FindById(el.ProjectId)
	author, _ := s.repo.Author.FindById(el.AuthorId)
	assignee, _ := s.repo.Author.FindById(el.AssigneeId)
	return dto.Issue{
		Id: el.ID,
		Project: dto.Project{
			Id:    project.Id,
			Title: project.Title,
			Key:   project.Key,
		},
		Key:         el.Key,
		CreatedTime: el.CreatedTime,
		ClosedTime:  el.ClosedTime,
		UpdatedTime: el.UpdatedTime,
		Summary:     el.Summary,
		Description: el.Description,
		Priority:    el.Priority,
		Creator:     author.Name,
		Assignee:    assignee.Name,
		Status:      el.Status,
	}
}
