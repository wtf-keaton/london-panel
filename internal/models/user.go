package models

import (
	"encoding/json"
	"gorm.io/datatypes"
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

	Cheats    datatypes.JSON
	CheatList CheatList `gorm:"-"`
}

func (u *UserModel) AfterFind(tx *gorm.DB) (err error) {
	json.Unmarshal(u.Cheats, &u.CheatList)
	return
}
