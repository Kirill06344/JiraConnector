package application

import (
	"github.com/stewie/config"
)

type Application struct {
	cfg *config.Config
}

var App = Configure()

func Configure() *Application {
	cfg, _ := config.Load()
	return &Application{cfg: cfg}
}

func (app *Application) Config() *config.Config {
	return app.cfg
}
