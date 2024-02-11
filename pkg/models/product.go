package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint   `gorm:"primarykey"`
	Uuid        string `gorm:"unique"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text;not null"`
	Price       int    `gorm:"not null;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func NewProduct() *Product {
	return &Product{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}
