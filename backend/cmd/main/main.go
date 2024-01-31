package main

import (
	"backend/internal/controller"
	"backend/internal/database"
	"backend/internal/router"
	"backend/internal/service"
	"backend/internal/utils"
	"fmt"
	_ "github.com/swaggo/swag"
	"net/http"
)

// @title JIRA_Analyzer
// @version	1.0
// @description	Разработка промышленного клиент-серверного приложения с применением принципов микросервисной архитектуры, языков программирования Golang, фрейморка Angular и TypeScript.
// @BasePath /
// @host localhost:8000

func main() {
	config, err := utils.GetConfig("config.yml")
	if err != nil {
		utils.Logger.Fatal(err)
		return
	}

	db, err := database.NewDB(config)
	if err != nil {
		utils.Logger.Fatal(err)
		return
	}

	serviceGroup := service.NewService(db.Repository())

	controllerGroup := controller.NewGroup(serviceGroup)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: router.NewRouter(controllerGroup),
	}

	err = server.ListenAndServe()

	utils.Logger.Fatal(err)
}
