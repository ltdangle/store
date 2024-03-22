package models

import (
	"store/pkg/i"
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	Uuid      string         `gorm:"primarykey" db:"uuid"`
	UserUuid  string         `db:"user_uuid"`
	Total     int            `gorm:"not null;default:0" db:"total"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at"`
}

func NewCart() *Cart {
	return &Cart{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}

func (cart Cart) PrimaryKey() string {
	return "uuid"
}

func (cart Cart) TableName() string {
	return "carts"
}

func (cart Cart) New() i.AdminEntity {
	return NewCart()
}

func (cart *Cart) GetUuid() string {
	return cart.Uuid
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

func (cartItem CartItem) PrimaryKey() string {
	return "uuid"
}

func (cartItem CartItem) TableName() string {
	return "cart_items"
}

func (cartItem CartItem) String() string {
	return cartItem.Uuid
}
func (cart CartItem) New() i.AdminEntity {
	return NewCartItem()
}
func (cartItem *CartItem) GetUuid() string {
	return cartItem.Uuid
}
