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

// GetAllProjects GetProjects retrieving all projects
// @Summary Retrieving all projects from database
// @Description Retrieving all projects from database
// @Success 200 {array} dto.Project
// @Tags project
// @Router /api/v1/projects [get]
func (pc *Project) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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

// GetProject retrieving project by id
// @Summary Get project by id
// @Description Get project by id
// @Tags project
// @Router /api/v1/projects/{id} [get]
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

// CreateProject Create new project
// @Summary Create new project
// @Description Create new project
// @Success 201
// @Tags project
// @Router /api/v1/projects [post]
func (pc *Project) CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
		Description: input.Title,
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

// UpdateProject  Update project
// @Summary Update project by id
// @Description Update project
// @Success 200
// @Tags project
// @Router /api/v1/projects [patch]
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
		Description: input.Title,
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

// DeleteProject  Delete project
// @Summary Delete project by id
// @Description Delete project
// @Success 200
// @Tags project
// @Router /api/v1/projects{id} [delete]
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
