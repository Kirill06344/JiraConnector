package database

import (
	"github.com/stewie/internal/dataTransformer"
	"github.com/stewie/internal/dto"
	"github.com/stewie/internal/entity"
	"gorm.io/gorm"
)

func (repo *DB) InsertData(project *dto.Project, issues []dto.Issue) error {
	db := repo.db
	err := db.Transaction(func(tx *gorm.DB) error {
		projectId, err := repo.saveProject(project)
		if err != nil {
			return err
		}

		issueEntities := make([]entity.Issue, len(issues))
		for i, issue := range issues {
			authorId, err := repo.saveAuthor(&issue.Fields.Creator)
			if err != nil {
				return err
			}

			assigneeId, err := repo.saveAuthor(&issue.Fields.Assignee)
			if err != nil {
				return err
			}

			issueEntities[i] = dataTransformer.IssueToEntity(&issue, projectId, authorId, assigneeId)
		}

		result := db.Create(&issueEntities)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
	return err
}

func (repo *DB) saveProject(project *dto.Project) (uint, error) {
	db := repo.db
	var projectEntity entity.Project
	result := db.Where("key=?", project.Key).Find(&projectEntity)
	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected != 0 {
		db.Delete(&projectEntity)
	}
	projectEntity = dataTransformer.ProjectToEntity(project)
	result = db.Create(&projectEntity)
	if result.Error != nil {
		return 0, result.Error
	}
	return projectEntity.ID, nil
}

func (repo *DB) saveAuthor(author *dto.Author) (uint, error) {
	db := repo.db
	var authorEntity entity.Author
	result := db.Where("name=?", author.DisplayName).Find(&authorEntity)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		authorEntity = dataTransformer.AuthorToEntity(author)
		result = db.Create(&authorEntity)
		if result.Error != nil {
			return 0, result.Error
		}
	}
	return authorEntity.ID, nil
}
