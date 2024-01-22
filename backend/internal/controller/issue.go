package controller

import (
	"backend/internal/service"
	"backend/internal/utils"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type Issue struct {
	service *service.IssueService
}

func NewIssueController(service *service.IssueService) *Issue {
	return &Issue{service: service}
}

func (ic *Issue) GetAllIssues(w http.ResponseWriter, r *http.Request) {
	name := "Jira Analyzer REST API Get Issues"
	issues, err := ic.service.Find()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ServerErrorResponse(w, r, err, name)
		return
	}
	response := utils.Envelope{
		"data":    issues,
		"message": "success",
		"name":    name,
		"status":  true,
	}

	err = utils.WriteJSON(w, http.StatusOK, response, nil)
	if err != nil {
		ServerErrorResponse(w, r, err, name)
	}
}

func (ic *Issue) GetIssue(w http.ResponseWriter, r *http.Request) {
	name := "Jira Analyzer REST API Get Issue"
	id, err := utils.ReadIdParam(r)
	if err != nil {
		BadRequestResponse(w, r, err, name)
		return
	}
	issue, err := ic.service.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			NotFoundResponse(w, r, name)
			return
		}
		ServerErrorResponse(w, r, err, name)
		return
	}
	response := utils.Envelope{
		"data":    issue,
		"message": "success",
		"name":    name,
		"status":  true,
	}

	err = utils.WriteJSON(w, http.StatusOK, response, nil)
	if err != nil {
		ServerErrorResponse(w, r, err, name)
	}
}

func (ic *Issue) CreateIssue(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (ic *Issue) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (ic *Issue) DeleteIssue(w http.ResponseWriter, r *http.Request) {
	name := "Jira Analyzer REST API Delete Issue"
	id, err := utils.ReadIdParam(r)
	if err != nil {
		BadRequestResponse(w, r, err, name)
		return
	}

	err = ic.service.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			NotFoundResponse(w, r, name)
			return
		}
		ServerErrorResponse(w, r, err, name)
		return
	}

	response := utils.Envelope{
		"message": "success",
		"name":    name,
		"status":  true,
	}

	err = utils.WriteJSON(w, http.StatusOK, response, nil)
	if err != nil {
		ServerErrorResponse(w, r, err, name)
	}
}
