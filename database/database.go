package database

import (
	"os"

	"github.com/salamanderman234/simple-api/config"
	model "github.com/salamanderman234/simple-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dbConn, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        os.Getenv(config.GetDatabaseUrl()),
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func ConnectAndMigrate() error {
	dbConn, err := Connect()
	if err != nil {
		return err
	}
	err = Migrate(dbConn)
	if err != nil {
		return err
	}
	model.SetDatabaseConnection(dbConn)
	return nil
}

func Migrate(connection *gorm.DB) error {
	return nil
}
