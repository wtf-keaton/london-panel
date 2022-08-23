package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model

	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}
