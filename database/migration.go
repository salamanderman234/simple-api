package database

import (
	model "github.com/salamanderman234/simple-api/models"
	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) error {
	err := conn.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		return err
	}
	return nil
}
