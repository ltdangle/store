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
	Uuid string `gorm:"primarykey" db:"uuid"`

	CartUuid    string         `db:"cart_uuid"`
	ProductUuid string         `db:"product_uuid"`
	Quantity    int            `gorm:"not null;default:1" db:"quantity"`
	Subtotal    int            `gorm:"not null;default:0" db:"subtotal"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
	DeletedAt   gorm.DeletedAt `db:"deleted_at"`
}

func NewCartItem() *CartItem {
	return &CartItem{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}

func (cartItem *CartItem) PrimaryKey() string {
	return "uuid"
}

func (cartItem *CartItem) TableName() string {
	return "cart_items"
}

func (cartItem *CartItem) String() string {
	return cartItem.Uuid
}
