package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(notFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowedResponse)

	resources(router)

	return router
}

func resources(router *mux.Router) {
	router.HandleFunc("/api/v1/issues", getAllIssues).Methods("GET")

	router.HandleFunc("/api/v1/issues/{id}", getIssue).Methods("GET")
	router.HandleFunc("/api/v1/issues/{id}", createIssue).Methods("POST")
	router.HandleFunc("/api/v1/issues/{id}", updateIssue).Methods("PUT")
	router.HandleFunc("/api/v1/issues/{id}", deleteIssue).Methods("DELETE")
}
