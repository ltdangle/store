package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Uuid           string
	ShoppingCartId uint `gorm:"not null;index"`
	ProductId      uint `gorm:"not null;index"`
	Quantity       int  `gorm:"not null;default:1"`
	Subtotal       int  `gorm:"not null;default:0"`
}
