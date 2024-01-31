package controller

import (
	"backend/internal/dto"
	"backend/internal/service"
	"backend/internal/utils"
	"net/http"
)

type Connector struct {
	service *service.ConnectorService
}

func NewConnectorController(service *service.ConnectorService) *Connector {
	return &Connector{
		service: service,
	}
}

// GetAllProjects GetProjects retrieving all projects
// @Summary Retrieving all projects
// @Description Retrieving all projects
// @Tags connector
// @Param limit query integer true "Projects count on one page"
// @Param page query integer true "Number of page"
// @Param search query string true "Search for project name"
// @Router /api/v1/connector/projects [get]
func (c *Connector) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	name := "Jira Analyzer REST API Get All projects"
	params, err := utils.GetQueryParams(r)
	if err != nil {
		BadRequestResponse(w, r, err, name)
		return
	}

	response, err := c.service.GetAllProjects(params)
	if err != nil {
		ServerErrorResponse(w, r, err, name)
		return
	}

	projects := make([]dto.Project, len(response.Projects))
	for i, el := range response.Projects {
		projects[i] = dto.Project{
			URL:         el.Url,
			Description: el.Description,
			Key:         el.Key,
			Name:        el.Name,
		}
	}

	env := utils.Envelope{
		"data":     projects,
		"message":  "success",
		"name":     name,
		"pageInfo": response.GetPageInfo(),
		"status":   true,
	}

	err = utils.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		ServerErrorResponse(w, r, err, name)
	}

}

// DownloadProject downloads project with given key from Jira
// @Summary Downloads project with given key from Jira
// @Description Downloads project with given key from Jira
// @Tags connector
// @Param key query string true "Project's key"
// @Success 200 {object} dto.Project
// @Failure 400 {string} string "bad request"
// @Router /updateProject [post]
func (c *Connector) DownloadProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	name := "Jira Analyzer REST API Download Project"

	key, err := utils.GetKeyQuery(r)
	if err != nil {
		BadRequestResponse(w, r, err, name)
		return
	}
	id, err := c.service.DownloadProject(key)
	env := utils.Envelope{
		"data":    dto.Project{Id: id},
		"message": "ok",
		"name":    name,
		"status":  "success",
	}
	err = utils.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		ServerErrorResponse(w, r, err, name)
	}
}
