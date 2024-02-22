package models

import (
	"store/pkg/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	cfg := infra.ReadConfig("../../.env")
	db := infra.Gorm(cfg.POSTGRES_URL)
	field1 := NewBaseProductField()
	field1.Type = "string"
	field1.Title = "Instructions"
	field1.Description = "Please enter your order instructions"
	field1.Value = "This is default value"

	product := NewBaseProduct()
	product.Fields = append(product.Fields, field1)

	// Save product.
	tx := db.Save(product)
	if tx.Error != nil {
		t.Fatal(tx.Error)
	}

	// Retrieve product with fields.
	var fndPrdct BaseProduct

	result := db.Preload("Fields").Where("uuid= ?", product.Uuid).First(&fndPrdct)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	assert.Equal(t, product.Uuid, fndPrdct.Uuid)
}
