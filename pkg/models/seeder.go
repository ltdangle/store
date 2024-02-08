package models

import (
	"math/rand"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}
func (s *Seeder) CreateUser() *User {
	customer := NewUser()
	customer.Email = faker.Name()
	s.db.Create(customer)
	return customer
}

func (s *Seeder) CreateCustomer(user *User) *Customer {
	customer := NewCustomer(user)
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

func (s *Seeder) CreateProduct() *Product {
	product := NewProduct()
	product.Name = s.generateRandomString(10)
	s.db.Create(product)
	return product
}

func (s *Seeder) AddProduct(cartItem *CartItem, product *Product) {
	cartItem.ProductUuid = product.Uuid
	s.db.Save(cartItem)
}

func (s *Seeder) generateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
