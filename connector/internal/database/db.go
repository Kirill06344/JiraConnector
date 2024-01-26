package database

import (
	"fmt"
	"github.com/stewie/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

func NewDB(cfg *config.Config) (*DB, error) {
	db, err := connectDB(cfg)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}

func connectDB(cfg *config.Config) (*gorm.DB, error) {
	dbCfg := cfg.DB
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbCfg.User, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	fmt.Println(dsn)

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
		//utils.Logger.Fatalln("Connection to database is not established")
		return nil, err
	}

	return db, err
}
