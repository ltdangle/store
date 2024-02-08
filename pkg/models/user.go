package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Uuid      string `gorm:"unique"`
	Email     string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func NewUser() *User {
	return &User{
		Uuid:      NewUuid(),
		CreatedAt: time.Now(),
	}
}
