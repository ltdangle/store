package models

import (
	"time"

	"gorm.io/gorm"
)

type ShoppingCart struct {
	ID        uint `gorm:"primarykey"`
	Uuid      string
	UserUuid  string `gorm:"not null;index"`
	Total     int    `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewCart() *ShoppingCart {
	return &ShoppingCart{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}
