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

func (seeder *Seeder) BuildFurnitureProduct(name string, description string) {

	instrctns := NewProductField()
	instrctns.Type = "string"
	instrctns.Title = "Instructions"
	instrctns.Description = "Order instructions"
	instrctns.Value = ""

	date := NewProductField()
	date.Type = "date"
	date.Title = "Delivery date"
	date.Description = "Delivery date of the product"
	date.Value = ""

	file := NewProductField()
	file.Type = "attachment"
	file.Title = "Attachment"
	file.Description = "Project attachment"
	file.Value = ""

	size := NewProductField()
	size.Type = "string"
	size.Title = "Size"
	size.Description = "Project size"
	size.Value = ""

	color := NewProductField()
	color.Type = "string"
	color.Title = "Color"
	color.Description = "Project color"
	color.Value = ""

	product := NewProduct()
	product.Type = "furniture"
	product.Name = name
	product.Description = description
	product.Fields = append(product.Fields, instrctns)
	product.Fields = append(product.Fields, date)
	product.Fields = append(product.Fields, file)
	product.Fields = append(product.Fields, size)
	product.Fields = append(product.Fields, color)

	// Save product.
	tx := seeder.db.Save(product)
	if tx.Error != nil {
		panic(tx.Error)
	}
}
