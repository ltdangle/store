package models

import (
	"gorm.io/gorm"
)


type Product struct {
	gorm.Model
	Uuid
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text;not null"`
	Price       int    `gorm:"not null;default:0"`
}

// ShoppingCart represents a user's shopping cart.
type ShoppingCart struct {
	gorm.Model
	Uuid
	UserId uint `gorm:"not null;index"`
	Total  int  `gorm:"not null;default:0"`
}

// CartItem represents a single line item in a shopping cart.
type CartItem struct {
	gorm.Model
	Uuid
	ShoppingCartId uint `gorm:"not null;index"`     // ForeignKey to ShoppingCart model
	ProductId      uint `gorm:"not null;index"`     // ForeignKey to the Product model
	Quantity       int  `gorm:"not null;default:1"` // Quantity of the product in the cart
	Subtotal       int  `gorm:"not null;default:0"` // Subtotal should be calculated based on Quantity * Product.Price
}
