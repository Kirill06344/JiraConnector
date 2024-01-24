package controller

import (
	"backend/internal/dto"
	"backend/internal/service"
	"backend/internal/utils"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type Project struct {
	service *service.ProjectService
}

func NewProjectController(service *service.ProjectService) *Project {
	return &Project{service: service}
}

func (pc *Project) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	name := "Jira Analyzer REST API Get Projects"
	projects, err := pc.service.Find()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ServerErrorResponse(w, r, err, name)
		return
	}
	response := utils.Envelope{
		"data":    projects,
		"message": "success",
		"name":    name,
		"status":  true,
	}

	err = utils.WriteJSON(w, http.StatusOK, response, nil)
	if err != nil {
		ServerErrorResponse(w, r, err, name)
	}
}

func (pc *Project) GetProject(w http.ResponseWriter, r *http.Request) {
	name := "Jira Analyzer REST API Get Project"
	id, err := utils.ReadIdParam(r)
	if err != nil {
		BadRequestResponse(w, r, err, name)
		return
	}
	project, err := pc.service.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			NotFoundResponse(w, r, name)
			return
		}
		ServerErrorResponse(w, r, err, name)
		return
	}
	response := utils.Envelope{
		"data":    project,
		"message": "success",
		"name":    name,
		"status":  true,
	}

	err = utils.WriteJSON(w, http.StatusOK, response, nil)
	if err != nil {
		ServerErrorResponse(w, r, err, name)
	}
}

func (pc *Project) CreateProject(w http.ResponseWriter, r *http.Request) {
	name := "Jira Analyzer REST API Create Project"
	var input struct {
		Title string `json:"Title"`
	}

	err := utils.ReadJSON(w, r, &input)
	if err != nil {
		BadRequestResponse(w, r, err, name)
		return
	}

	project := &dto.Project{
		Title: input.Title,
	}

	err = pc.service.Create(project)
	if err != nil {
		ServerErrorResponse(w, r, err, name)
	}

	response := utils.Envelope{
		"message": "success",
		"name":    name,
		"status":  true,
	}

	err = utils.WriteJSON(w, http.StatusCreated, response, nil)
	if err != nil {
		ServerErrorResponse(w, r, err, name)
	}
}

func (pc *Project) UpdateProject(w http.ResponseWriter, r *http.Request) {
	name := "Jira Analyzer REST API Update Project"
	id, err := utils.ReadIdParam(r)
	if err != nil {
		BadRequestResponse(w, r, err, name)
		return
	}

	var input struct {
		Title string `json:"Title"`
	}

	err = utils.ReadJSON(w, r, &input)
	if err != nil {
		BadRequestResponse(w, r, err, name)
		return
	}

	project := &dto.Project{
		Title: input.Title,
	}

	err = pc.service.Update(id, project)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			NotFoundResponse(w, r, name)
		default:
			ServerErrorResponse(w, r, err, name)
		}
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

func (pc *Project) DeleteProject(w http.ResponseWriter, r *http.Request) {
	name := "Jira Analyzer REST API Delete Project"
	id, err := utils.ReadIdParam(r)
	if err != nil {
		BadRequestResponse(w, r, err, name)
		return
	}

	err = pc.service.Delete(id)
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
