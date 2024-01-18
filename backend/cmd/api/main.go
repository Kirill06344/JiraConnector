package main

import (
	"backend/internal"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type application struct {
	cfg    *internal.Config
	logger *logrus.Logger
}

func main() {
	log, err := internal.ConfigureLogger()
	if err != nil {
		log.Fatal(err)
		return
	}

	config, err := internal.GetConfig("config.yml")
	if err != nil {
		log.Fatal(err)
		return
	}

	app := &application{
		cfg:    config,
		logger: log,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: app.routes(),
	}

	err = server.ListenAndServe()

	log.Fatal(err)
}
