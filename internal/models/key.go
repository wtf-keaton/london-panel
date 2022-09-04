package models

import (
	"gorm.io/gorm"
	"time"
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
	Hours   int
	EndTime time.Time

	// User information
	CreatedBy string
}

type LogModel struct {
	gorm.Model

	// Basic Information
	KeyID   int
	Message string
}
