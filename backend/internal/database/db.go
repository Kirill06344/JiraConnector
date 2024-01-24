package database

import (
	"backend/internal/repository"
	"backend/internal/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DB struct {
	db         *gorm.DB
	repository *repository.Repository
}

func (db *DB) Repository() *repository.Repository {
	return db.repository
}

func (db *DB) IssueRepository() repository.IssueRepository {
	return db.repository.Issue
}

func (db *DB) ProjectRepository() repository.ProjectRepository {
	return db.repository.Project
}

func NewDB(cfg *utils.Config) (*DB, error) {
	db, err := connectDB(cfg)
	if err != nil {
		return nil, err
	}
	return &DB{
		db:         db,
		repository: repository.NewRepository(db),
	}, nil
}

func connectDB(cfg *utils.Config) (*gorm.DB, error) {
	dbCfg := cfg.Database
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	utils.Logger.Print("Open initialize db session")

	gormLogger := logger.New(
		utils.Logger,
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
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
