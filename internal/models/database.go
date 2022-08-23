package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"template/pgk/settings"
)

var DB *gorm.DB

func Connect() {
	DB, _ = gorm.Open(
		mysql.Open(settings.Settings.DSN),
		&gorm.Config{},
	)

	defer makeMigrations()
}

func makeMigrations() {
	err := DB.AutoMigrate(
		&UserModel{},
		&ConfigModel{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}
