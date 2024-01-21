package router

import (
	"backend/internal/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(controllers *controller.Group) *mux.Router {
	router := mux.NewRouter()

	//router.NotFoundHandler = http.HandlerFunc(controller.NotFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(controller.MethodNotAllowedResponse)

	setProjectPaths(controllers.Project, router)

	return router
}

func setProjectPaths(pc *controller.Project, router *mux.Router) {
	router.HandleFunc("/api/v1/projects", pc.GetAllProjects).Methods("GET")

	router.HandleFunc("/api/v1/projects/{id}", pc.GetProject).Methods("GET")
	router.HandleFunc("/api/v1/projects", pc.CreateProject).Methods("POST")
	router.HandleFunc("/api/v1/projects/{id}", pc.UpdateProject).Methods("PUT")
	router.HandleFunc("/api/v1/projects/{id}", pc.DeleteProject).Methods("DELETE")
}
