package db

import (
	"fmt"
	"time"

	usersModels "github.com/thmelodev/ddd-events-api/src/modules/auth/infra/models"
	eventsModel "github.com/thmelodev/ddd-events-api/src/modules/events/infra/models"
	"github.com/thmelodev/ddd-events-api/src/providers/config"
	"github.com/thmelodev/ddd-events-api/src/utils/hash"
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

	db.AutoMigrate(&eventsModel.EventModel{}, &usersModels.UserModel{})

	createAdminUser(db, config)

	return &GormDatabase{
		DB: db,
	}, nil
}

func createAdminUser(db *gorm.DB, config *config.Config) {
	var admin usersModels.UserModel
	result := db.First(&admin, "email = ?", "admin@teste.com")

	hashedPassword, err := hash.HashPassword("123456")

	if err != nil {
		panic(fmt.Errorf("failed to hash admin password: %v", err))
	}

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		admin = usersModels.UserModel{
			Id:       config.App.AdminId,
			Email:    "admin@teste.com",
			Password: hashedPassword,
		}

		err = db.Create(&admin).Error

		if err != nil {
			panic(fmt.Errorf("failed to create admin user: %v", err))
		} else {
			fmt.Println("Admin user created successfully.")
		}
	} else if result.Error != nil {
		panic(fmt.Errorf("error checking admin user: %v", result.Error))
	} else {
		fmt.Println("Admin user already exists.")
	}
}
