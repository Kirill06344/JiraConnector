package application

import (
	"github.com/stewie/config"
	"github.com/stewie/internal/database"
)

type Application struct {
	cfg *config.Config
	db  *database.DB
}

var App *Application

func Configure() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := database.NewDB(cfg)
	if err != nil {
		return err
	}

	App = &Application{
		cfg: cfg,
		db:  db,
	}
	return nil
}

func (app *Application) Config() *config.Config {
	return app.cfg
}

func (app *Application) DB() *database.DB {
	return app.db
}
