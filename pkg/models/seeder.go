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

func (seeder *Seeder) BuildBasicProduct(name string, description string) *BaseProduct {

	instrctns := NewBaseProductField()
	instrctns.Type = "string"
	instrctns.Title = "Instructions"
	instrctns.Description = "Order instructions"
	instrctns.Value = "Please take your time and do a good job. Quality over delivery time is preferred."

	date := NewBaseProductField()
	date.Type = "date"
	date.Title = "Delivery date"
	date.Description = "Delivery date of the product"
	date.Value = "21/11/2024"

	file := NewBaseProductField()
	file.Type = "attachment"
	file.Title = "Attachment"
	file.Description = "Project attachment"
	file.Value = "file1.jpg"

	size := NewBaseProductField()
	size.Type = "string"
	size.Title = "Size"
	size.Description = "Project size"
	size.Value = "XL"

	color := NewBaseProductField()
	color.Type = "string"
	color.Title = "Color"
	color.Description = "Project color"
	color.Value = "beige"

	product := NewBaseProduct()
	product.Type = "furniture"
	product.Name = name
	product.Description = description
	// TODO: 
	// product.Fields = append(product.Fields, instrctns)
	// product.Fields = append(product.Fields, date)
	// product.Fields = append(product.Fields, file)
	// product.Fields = append(product.Fields, size)
	// product.Fields = append(product.Fields, color)

	// Save product.
	tx := seeder.db.Save(product)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return product
}
