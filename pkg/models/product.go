package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID   uint   `gorm:"primarykey"`
	Uuid string `gorm:"unique"`

	Type        string
	Name        string          `gorm:"type:varchar(255);not null"`
	Description string          `gorm:"type:text;not null"`
	BasePrice   int             `gorm:"not null;default:0"`
	Fields      []*ProductField `gorm:"foreignKey:ProductID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewProduct() *Product {
	return &Product{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}

type ProductField struct {
	ID          uint   `gorm:"primarykey"`
	Uuid        string `gorm:"unique"`
	Type        string
	Title       string
	Description string
	Value       string

	ProductID uint

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewProductField() *ProductField {
	return &ProductField{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}
