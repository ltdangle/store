package models

import (
	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

func (seeder *Seeder) BuildBasicProduct(name string, description string) *Product {

	product := NewProduct()
	product.Type = "furniture"
	product.Name = name
	product.Description = description

	// Save product.
	tx := seeder.db.Save(product)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return product
}
