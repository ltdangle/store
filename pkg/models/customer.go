package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Uuid
	Email string `gorm:"type:varchar(255);not null"`
}

func NewCustomer() *Customer {
	return &Customer{Uuid: NewUuid()}
}
