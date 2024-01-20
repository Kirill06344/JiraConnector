package main

import (
	"backend/internal/controller"
	"backend/internal/router"
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

	issue := &controller.IssueController{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: router.NewRouter(issue),
	}

	err = server.ListenAndServe()

	utils.Logger.Fatal(err)
}
