package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseProduct struct {
	Uuid string `gorm:"primarykey"`

	Type        string
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text;not null"`
	BasePrice   int    `gorm:"not null;default:0"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewBaseProduct() *BaseProduct {
	return &BaseProduct{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}

type BaseProductField struct {
	Uuid string `gorm:"primarykey"`

	Type            string
	Title           string
	Description     string
	Value           string
	BaseProductUuid string 

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewBaseProductField() *BaseProductField {
	return &BaseProductField{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}
