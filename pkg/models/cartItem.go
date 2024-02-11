package models

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	ID       uint   `gorm:"primarykey"`
	Uuid     string `gorm:"unique"`
	CartUuid string

	CartID uint

	Quantity  int `gorm:"not null;default:1"`
	Subtotal  int `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewCartItem() *CartItem {
	return &CartItem{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}
