package dataTransformer

import (
	"github.com/stewie/internal/dto"
	"github.com/stewie/internal/entity"
	"time"
)

var ISO8061 = "2006-01-02T15:04:05.999+0000"

func ProjectToEntity(project *dto.Project) entity.Project {
	return entity.Project{
		Title: project.Description,
		Key:   project.Key,
		Name:  project.Name,
		Url:   project.URL,
	}
}

func AuthorToEntity(author *dto.Author) entity.Author {
	return entity.Author{
		Name: author.DisplayName,
	}
}

func IssueToEntity(issue *dto.Issue, projectId uint, authorId uint, assigneeId uint) entity.Issue {
	createdTime, _ := time.Parse(ISO8061, issue.Fields.CreatedTime)
	updatedTime, _ := time.Parse(ISO8061, issue.Fields.UpdatedTime)

	return entity.Issue{
		ProjectId:   projectId,
		AuthorId:    authorId,
		AssigneeId:  assigneeId,
		Key:         issue.Key,
		CreatedTime: createdTime,
		UpdatedTime: updatedTime,
		Summary:     issue.Fields.Summary,
		Description: issue.Fields.Description,
		Priority:    issue.Fields.Priority.Name,
		Status:      issue.Fields.Status.Name,
		Type:        issue.Fields.Type.Name,
	}
}
