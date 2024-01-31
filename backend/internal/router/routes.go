package router

import (
	_ "backend/docs"
	"backend/internal/controller"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func NewRouter(controllers *controller.Group) *mux.Router {
	router := mux.NewRouter()

	//router.NotFoundHandler = http.HandlerFunc(controller.NotFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(controller.MethodNotAllowedResponse)

	setProjectPaths(controllers.Project, router)
	setIssuesPaths(controllers.Issue, router)
	setConnectorPaths(controllers.Connector, router)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return router
}

func setProjectPaths(pc *controller.Project, router *mux.Router) {
	router.HandleFunc("/api/v1/projects", pc.GetAllProjects).Methods("GET")
	router.HandleFunc("/api/v1/projects/{id}", pc.GetProject).Methods("GET")
	router.HandleFunc("/api/v1/projects", pc.CreateProject).Methods("POST")
	router.HandleFunc("/api/v1/projects/{id}", pc.UpdateProject).Methods("PUT")
	router.HandleFunc("/api/v1/projects/{id}", pc.DeleteProject).Methods("DELETE")
}

func setIssuesPaths(ic *controller.Issue, router *mux.Router) {
	router.HandleFunc("/api/v1/issues", ic.GetAllIssues).Methods("GET")
	router.HandleFunc("/api/v1/issues/{id}", ic.GetIssue).Methods("GET")
	router.HandleFunc("/api/v1/issues", ic.CreateIssue).Methods("POST")
	router.HandleFunc("/api/v1/issues/{id}", ic.UpdateIssue).Methods("PUT")
	router.HandleFunc("/api/v1/issues/{id}", ic.DeleteIssue).Methods("DELETE")
}

func setConnectorPaths(cc *controller.Connector, router *mux.Router) {
	router.HandleFunc("/api/v1/connector/projects", cc.GetAllProjects).Methods("GET")
	router.HandleFunc("/api/v1/connector/updateProject", cc.DownloadProject).Methods("GET")
}
