package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint `gorm:"primarykey"`
	Uuid      string
	UserUuid  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewCustomer(user *User) *Customer {
	return &Customer{
		Uuid:      NewUuid(),
		UserUuid:  user.Uuid,
		CreatedAt: time.Now(),
	}
}
