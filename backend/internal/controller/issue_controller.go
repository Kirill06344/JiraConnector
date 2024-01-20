package controller

import (
	response2 "backend/internal/dto/response"
	utils "backend/internal/utils"
	"net/http"
	"time"
)

type IssueController struct {
}

func (ic *IssueController) GetAllIssues(w http.ResponseWriter, r *http.Request) {

}

func (ic *IssueController) GetIssue(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIdParam(r)
	if err != nil {
		NotFoundResponse(w, r)
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
		ServerErrorResponse(w, r, err)
	}
}

func (ic *IssueController) CreateIssue(w http.ResponseWriter, r *http.Request) {
	//need to change this and declare this struct in any file
	var input struct {
		Id  int64  `json:"Id"`
		Key string `json:"Key"`
	}

	err := utils.ReadJSON(w, r, &input)
	if err != nil {
		BadRequestResponse(w, r, err)
		return
	}

}

func (ic *IssueController) UpdateIssue(w http.ResponseWriter, r *http.Request) {

}

func (ic *IssueController) DeleteIssue(w http.ResponseWriter, r *http.Request) {

}

func tmpIssue(id int64) response2.Issue {
	return response2.Issue{
		Id:          id,
		Project:     response2.Project{Id: 123, Title: "asdasd"},
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
