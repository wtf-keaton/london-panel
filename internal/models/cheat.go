package models

import "gorm.io/gorm"

type CheatModel struct {
	gorm.Model

	// General information
	Name    string
	Status  int
	Creator string

	// Developer settings
	Filename string
	Process  string
	Anticheat string
}
