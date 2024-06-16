package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID 		 uint
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	Role     string
	IsActive bool   `gorm:"default:true"`
}
