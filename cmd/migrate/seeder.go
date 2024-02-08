package main

import (
	"store/pkg/models"

	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}
func (s *Seeder) CreateCustomer() {
}
func (s *Seeder) AddCart(customer models.Customer) {

}
