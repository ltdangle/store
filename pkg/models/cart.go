package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint   `gorm:"primarykey"`
	Uuid      string `gorm:"unique"`
	UserUuid  string
	CartItems []*CartItem `gorm:"foreignKey:CartID"`
	Total     int         `gorm:"not null;default:0"`
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
	ID       uint   `gorm:"primarykey"`
	Uuid     string `gorm:"unique"`
	CartUuid string

	CartID  uint
	Product Product `gorm:"foreignKey:CartItemID;references:ID"`

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
