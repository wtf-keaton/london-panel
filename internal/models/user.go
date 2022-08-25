package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model

	Username string `gorm:"unique"`
	Password string
	Status   string
}
