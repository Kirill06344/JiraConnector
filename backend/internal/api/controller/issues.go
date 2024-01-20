package controller

import (
	response "backend/internal/api/dto/response"
	utils "backend/internal/utils"
	"net/http"
	"time"
)

func getAllIssues(w http.ResponseWriter, r *http.Request) {

}

func getIssue(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIdParam(r)
	if err != nil {
		notFoundResponse(w, r)
		return
	}
	//....
	issue := tmpIssue(id)
	env := utils.Envelope{
		"data":    issue,
		"message": "success",
		"name":    "Jira Analyzer REST API Get Issue",
		"status":  true,
	}

	err = utils.WriteJSON(w, http.StatusOK, env, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}

func createIssue(w http.ResponseWriter, r *http.Request) {
	//need to change this and declare this struct in any file
	var input struct {
		Id  int64  `json:"Id"`
		Key string `json:"Key"`
	}

	err := utils.ReadJSON(w, r, &input)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

}

func updateIssue(w http.ResponseWriter, r *http.Request) {

}

func deleteIssue(w http.ResponseWriter, r *http.Request) {

}

func tmpIssue(id int64) response.Issue {
	return response.Issue{
		Id:          id,
		Project:     response.Project{Id: 123, Title: "asdasd"},
		Key:         "ddsd",
		CreatedTime: time.Time{},
		ClosedTime:  time.Time{},
		UpdatedTime: time.Time{},
		Summary:     "ss",
		Description: "ff",
		Priority:    "gg",
		Status:      "gg",
	}
}
