package models

import (
	"gorm.io/gorm"
)

type KeyModel struct {
	gorm.Model

	// Basic Information
	Keycode    string `gorm:"unique"`
	Cheat      string
	Status     int
	HardwareID string
	Banned     bool

	// Time Management
	Hours   int64
	EndTime int64

	// User information
	CreatedBy string
}

type LogModel struct {
	gorm.Model

	// Basic Information
	KeyID   int
	Message string
}
