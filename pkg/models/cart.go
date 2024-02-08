package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint `gorm:"primarykey"`
	Uuid      string
	UserUuid  string `gorm:"not null;index"`
	Total     int    `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewCart() *Cart {
	return &Cart{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}
