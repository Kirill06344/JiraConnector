package router

import (
	"backend/internal/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(pc *controller.Project) *mux.Router {
	router := mux.NewRouter()

	//router.NotFoundHandler = http.HandlerFunc(controller.NotFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(controller.MethodNotAllowedResponse)

	resources(pc, router)

	return router
}

func resources(pc *controller.Project, router *mux.Router) {
	router.HandleFunc("/api/v1/projects", pc.GetAllProjects).Methods("GET")

	router.HandleFunc("/api/v1/projects/{id}", pc.GetProject).Methods("GET")
	router.HandleFunc("/api/v1/projects/{id}", pc.CreateProject).Methods("POST")
	router.HandleFunc("/api/v1/projects/{id}", pc.UpdateProject).Methods("PUT")
	router.HandleFunc("/api/v1/projects/{id}", pc.DeleteProject).Methods("DELETE")
}
