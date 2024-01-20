package main

import (
	"backend/internal/controller"
	"backend/internal/repository"
	"backend/internal/router"
	"backend/internal/service"
	"backend/internal/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	config, err := utils.GetConfig("config.yml")
	if err != nil {
		utils.Logger.Fatal(err)
		return
	}

	db, err := connectDB(config)
	if err != nil {
		return
	}

	projectService := service.NewProjectService(repository.NewProjectRepository(db))

	projectController := controller.NewProjectController(&projectService)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: router.NewRouter(projectController),
	}

	err = server.ListenAndServe()

	utils.Logger.Fatal(err)
}

func connectDB(cfg *utils.Config) (*gorm.DB, error) {
	dbCfg := cfg.Database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	utils.Logger.Print("Open initialize db session")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	postgresDB, _ := db.DB()
	if err != nil {
		return nil, err
	}

	err = postgresDB.Ping()

	if err != nil {
		utils.Logger.Fatalln("Connection to database is not established")
		return nil, err
	}
	utils.Logger.Print("Connection to database is established")

	return db, err
}
