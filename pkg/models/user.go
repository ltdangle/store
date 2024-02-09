package models

import (
	"time"

	"gorm.io/gorm"
)

const USER_CUSTOMER = "customer"

type User struct {
	ID        uint   `gorm:"primarykey"`
	Uuid      string `gorm:"unique"`
	Type      string
	Email     string `gorm:"type:varchar(255);not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewCustomer() *User {
	return &User{
		Uuid:      NewUuid(),
		Type:      USER_CUSTOMER,
		CreatedAt: time.Now(),
	}
}
