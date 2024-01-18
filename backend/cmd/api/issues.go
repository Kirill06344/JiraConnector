package main

import (
	"backend/internal/data"
	"backend/utils"
	"net/http"
	"time"
)

func (app *application) getAllIssuesHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) getIssueHandler(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createIssueHandler(w http.ResponseWriter, r *http.Request) {
	//need to change this and declare this struct in any file
	var input struct {
		Id  int64  `json:"Id"`
		Key string `json:"Key"`
	}

	err := utils.ReadJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	app.logger.Printf("%+v\n", input)
}

func (app *application) updateIssueHandler(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteIssueHandler(w http.ResponseWriter, r *http.Request) {

}

func tmpIssue(id int64) data.Issue {
	return data.Issue{
		Id:          id,
		Project:     data.Project{Id: 123, Title: "asdasd"},
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
