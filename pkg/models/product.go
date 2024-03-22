package models

import (
	"store/pkg/i"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Uuid string `gorm:"primarykey" db:"uuid"`

	Type        string `db:"type"`
	Name        string `gorm:"type:varchar(255);not null" db:"name"`
	Description string `gorm:"type:text;not null" db:"description"`
	BasePrice   int    `gorm:"not null;default:0" db:"base_price"`

	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
	DeletedAt gorm.DeletedAt `db:"deleted_at"`
}

func NewProduct() *Product {
	return &Product{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}
func (product Product) PrimaryKey() string {
	return "uuid"
}

func (product Product) TableName() string {
	return "products"
}
func (cart Product) New() i.AdminEntity {
	return NewProduct()
}
