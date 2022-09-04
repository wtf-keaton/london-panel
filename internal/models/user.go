package models

import (
	"gorm.io/gorm"
)

type CheatList struct {
	Cheats []string
}

type UserModel struct {
	gorm.Model

	Username string `gorm:"unique"`
	Password string
	Status   string
}
