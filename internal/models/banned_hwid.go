package models

import "gorm.io/gorm"

type BannedHardware struct {
	gorm.Model

	HardwareID string
	Reason     string
}
