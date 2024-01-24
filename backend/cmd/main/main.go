package main

import (
	"backend/internal/controller"
	"backend/internal/database"
	"backend/internal/router"
	"backend/internal/service"
	"backend/internal/utils"
	"fmt"
	"net/http"
)

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
