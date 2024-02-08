package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Uuid        string
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text;not null"`
	Price       int    `gorm:"not null;default:0"`
}
