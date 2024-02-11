package models

import (
	"store/pkg/infra"
	"testing"
)

func TestProduct(t *testing.T) {
	cfg := infra.ReadConfig("../../.env")
	db := infra.Gorm(cfg)
	field1 := &ProductField{
		Type:        "string",
		Title:       "Instructions",
		Description: "Please enter your order instructions",
		Value:       "This is default value",
	}
	product := NewProduct()
	product.Fields = append(product.Fields, field1)
	db.Save(product)
	t.Log("Product saved!")
}
