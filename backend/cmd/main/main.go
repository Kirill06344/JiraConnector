package main

import (
	"backend/internal/api/controller"
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

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: controller.Routes(),
	}

	err = server.ListenAndServe()

	utils.Logger.Fatal(err)
}
