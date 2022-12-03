package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Age     int    `json:"age"`
}

func (u *User) CreateMe() error {
	result := Database.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
