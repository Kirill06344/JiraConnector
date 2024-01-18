package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowedResponse)

	app.resources(router)

	return router
}

func (app *application) resources(router *mux.Router) {
	router.HandleFunc("/api/v1/issues", app.getAllIssuesHandler).Methods("GET")

	router.HandleFunc("/api/v1/issues/{id}", app.getIssueHandler).Methods("GET")
	router.HandleFunc("/api/v1/issues/{id}", app.createIssueHandler).Methods("POST")
	router.HandleFunc("/api/v1/issues/{id}", app.updateIssueHandler).Methods("PUT")
	router.HandleFunc("/api/v1/issues/{id}", app.deleteIssueHandler).Methods("DELETE")
}
