package db

import (
	"fmt"
	"time"

	"github.com/thmelodev/ddd-events-api/src/modules/events/infra/models"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/utils/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type GormDatabase struct {
	DB *gorm.DB
}

func NewDatabase(config *config.Config) (*GormDatabase, error) {

	log := logger.Get()

	dbConnectionRetries := config.Db.ConnectionRetries
	if dbConnectionRetries == 0 {
		dbConnectionRetries = 3 // default value if undefined
	}

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=America/Sao_Paulo", config.Db.Host, config.Db.Port, config.Db.User, config.Db.Name, config.Db.Pass)
	var db *gorm.DB
	var err error

	newLogger := gormLogger.New(
		log,
		gormLogger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  gormLogger.Error, // Log everything
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	for i := 0; i < dbConnectionRetries; i++ {
		db, err = gorm.Open(postgres.Open(dbUri), &gorm.Config{
			Logger: newLogger,
		})

		if err == nil {
			break
		}

		time.Sleep(5 * time.Second)
	}

	if err != nil {
		errorMessage := fmt.Errorf("failed to connect to database, error: %s", err)
		panic(errorMessage)
	}

	sqlDb, err := db.DB()
	if err != nil {
		errorMessage := fmt.Errorf("failed to get sql.DB from gorm.DB, error: %s", err)
		panic(errorMessage)
	}

	sqlDb.SetMaxIdleConns(config.Db.MaxIdleConnections)
	sqlDb.SetMaxOpenConns(config.Db.MaxOpenConnections)
	sqlDb.SetConnMaxLifetime(config.Db.ConnectionMaxLifetime)

	db.AutoMigrate(&models.EventModel{})

	return &GormDatabase{
		DB: db,
	}, nil
}
