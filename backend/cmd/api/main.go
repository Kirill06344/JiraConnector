package main

import (
	"backend/cmd/internal"
	"fmt"
	"net/http"
)

type application struct {
}

func main() {
	app := &application{}
	logger, err := internal.ConfigureLogger()
	if err != nil {
		logger.Fatal(err)
		return
	}

	cfg, err := internal.GetConfig("config.yml")
	if err != nil {
		logger.Fatal(err)
		return
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: app.routes(),
	}

	err = server.ListenAndServe()

	logger.Fatal(err)
}
