package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Uuid string
	Email string `gorm:"type:varchar(255);not null"`
}

func NewCustomer() *Customer {
	return &Customer{
		Uuid:  NewUuid(),
		Model: gorm.Model{CreatedAt: time.Now()},
	}
}
