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
	//utils.Logger.Print("Open initialize db session")

	//gormLogger := logger.New(
	//	utils.Logger,
	//	logger.Config{
	//		SlowThreshold:             time.Second, // Slow SQL threshold
	//		LogLevel:                  logger.Info, // Log level
	//		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	//		ParameterizedQueries:      true,        // Don't include params in the SQL log
	//		Colorful:                  true,        // Disable color
	//	},
	//)

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
	//utils.Logger.Print("Connection to database is established")

	return db, err
}
