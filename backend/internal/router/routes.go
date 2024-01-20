package router

import (
	"backend/internal/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(issue *controller.IssueController) *mux.Router {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(controller.NotFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(controller.MethodNotAllowedResponse)

	resources(issue, router)

	return router
}

func resources(issue *controller.IssueController, router *mux.Router) {
	router.HandleFunc("/api/v1/issues", issue.GetAllIssues).Methods("GET")

	router.HandleFunc("/api/v1/issues/{id}", issue.GetIssue).Methods("GET")
	router.HandleFunc("/api/v1/issues/{id}", issue.CreateIssue).Methods("POST")
	router.HandleFunc("/api/v1/issues/{id}", issue.UpdateIssue).Methods("PUT")
	router.HandleFunc("/api/v1/issues/{id}", issue.DeleteIssue).Methods("DELETE")
}
