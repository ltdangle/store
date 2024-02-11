package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseProduct struct {
	ID   uint   `gorm:"primarykey"`
	Uuid string `gorm:"unique"`

	Type        string
	Name        string              `gorm:"type:varchar(255);not null"`
	Description string              `gorm:"type:text;not null"`
	BasePrice   int                 `gorm:"not null;default:0"`
	Fields      []*BaseProductField `gorm:"foreignKey:BaseProductID"`

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
	ID          uint   `gorm:"primarykey"`
	Uuid        string `gorm:"unique"`
	Type        string
	Title       string
	Description string
	Value       string

	BaseProductID uint

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
