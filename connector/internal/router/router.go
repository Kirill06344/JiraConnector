package router

import (
	"github.com/gorilla/mux"
	"github.com/stewie/internal/controller"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/projects", controller.GetProjects).Methods("GET")

	return router
}
