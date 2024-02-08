package models

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	ID          uint `gorm:"primarykey"`
	Uuid        string
	CartUuid    string `gorm:"not null;index"`
	ProductUuid string `gorm:"not null;index"`
	Quantity    int    `gorm:"not null;default:1"`
	Subtotal    int    `gorm:"not null;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func NewCartItem() *CartItem {
	return &CartItem{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}
