package models

import (
	"time"

	"gorm.io/gorm"
)

type ShoppingCart struct {
	gorm.Model
	Uuid     string
	UserUuid string `gorm:"not null;index"`
	Total    int    `gorm:"not null;default:0"`
}

func NewCart() *ShoppingCart {
	return &ShoppingCart{
		Uuid:  NewUuid(),
		Model: gorm.Model{CreatedAt: time.Now()},
	}
}
