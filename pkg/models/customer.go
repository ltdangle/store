package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint `gorm:"primarykey"`
	Uuid      string
	Email     string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewCustomer() *Customer {
	return &Customer{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}
