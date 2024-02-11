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

	instrctns := NewProductField(NewBaseProductField())
	instrctns.Type = "string"
	instrctns.Title = "Instructions"
	instrctns.Description = "Order instructions"
	instrctns.Value = ""

	date := NewProductField(NewBaseProductField())
	date.Type = "date"
	date.Title = "Delivery date"
	date.Description = "Delivery date of the product"
	date.Value = ""

	file := NewProductField(NewBaseProductField())
	file.Type = "attachment"
	file.Title = "Attachment"
	file.Description = "Project attachment"
	file.Value = ""

	size := NewProductField(NewBaseProductField())
	size.Type = "string"
	size.Title = "Size"
	size.Description = "Project size"
	size.Value = ""

	color := NewProductField(NewBaseProductField())
	color.Type = "string"
	color.Title = "Color"
	color.Description = "Project color"
	color.Value = ""

	product := NewProduct(NewBaseProduct())
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

func (seeder *Seeder) BuildBasicFurnitureProduct(name string, description string) {

	instrctns := NewBaseProductField()
	instrctns.Type = "string"
	instrctns.Title = "Instructions"
	instrctns.Description = "Order instructions"
	instrctns.Value = ""

	date := NewBaseProductField()
	date.Type = "date"
	date.Title = "Delivery date"
	date.Description = "Delivery date of the product"
	date.Value = ""

	file := NewBaseProductField()
	file.Type = "attachment"
	file.Title = "Attachment"
	file.Description = "Project attachment"
	file.Value = ""

	size := NewBaseProductField()
	size.Type = "string"
	size.Title = "Size"
	size.Description = "Project size"
	size.Value = ""

	color := NewBaseProductField()
	color.Type = "string"
	color.Title = "Color"
	color.Description = "Project color"
	color.Value = ""

	product := NewBaseProduct()
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
