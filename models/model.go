package model

import "gorm.io/gorm"

var Database *gorm.DB

type Model interface {
	CreateMe() error
	Update([]string, []interface{}) error
	DeleteMe(*gorm.DB) error
	Delete([]string, []interface{}) error
}

func SetDatabaseConnection(conn *gorm.DB) {
	Database = conn
}
