package models

import (
	"time"
)

type CartItem struct {
	ID             uint `gorm:"primarykey"`
	Uuid           string
	ShoppingCartId uint `gorm:"not null;index"`
	ProductId      uint `gorm:"not null;index"`
	Quantity       int  `gorm:"not null;default:1"`
	Subtotal       int  `gorm:"not null;default:0"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt time.Time
}
