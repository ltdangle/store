package models

import (
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}
func (s *Seeder) CreateCustomer() *Customer {
	customer := NewCustomer()
	customer.Email = faker.Name()
	s.db.Create(customer)
	return customer
}

func (s *Seeder) AddCart(customer *Customer) *Cart {
	cart := NewCart()
	cart.UserUuid = customer.Uuid
	s.db.Create(cart)
	return cart
}
func (s *Seeder) AddCartItem(cart *Cart) *CartItem {
	item := NewCartItem()
	item.CartUuid = cart.Uuid
	s.db.Create(item)
	return item
}
