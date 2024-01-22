package controller

import (
	"github.com/stewie/internal/connector"
	"github.com/stewie/internal/utils"
	"net/http"
)

func GetProjects(w http.ResponseWriter, r *http.Request) {
	params, err := utils.GetQueryParams(r)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}
	result, err := connector.GetProjects(params)
	err = utils.WriteJSON(w, http.StatusOK, result, nil)
	if err != nil {
		ServerErrorResponse(w, r, err)
	}
}
