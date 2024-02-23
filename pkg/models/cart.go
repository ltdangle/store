package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	Uuid      string `gorm:"primarykey"`
	UserUuid  string
	Total     int `gorm:"not null;default:0"`
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

type CartItem struct {
	Uuid     string `gorm:"primarykey"`

	CartUuid string
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
func (cartItem CartItem) String() string {
	return cartItem.Uuid
}
