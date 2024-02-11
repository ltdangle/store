package models

import (
	"fmt"
	"store/pkg/infra"
	"testing"
)

func TestProduct(t *testing.T) {
	cfg := infra.ReadConfig("../../.env")
	db := infra.Gorm(cfg)
	field1 := NewProductField()
	field1.Type = "string"
	field1.Title = "Instructions"
	field1.Description = "Please enter your order instructions"
	field1.Value = "This is default value"

	product := NewProduct()
	product.Fields = append(product.Fields, field1)

	tx := db.Save(product)
	if tx.Error != nil {
		panic(tx.Error)
	}
	t.Log("Product saved!")

	var fndPrdct Product

	result := db.Preload("Fields").Where("uuid= ?", product.Uuid).First(&fndPrdct)
	if result.Error != nil {
		t.Log(result.Error)
		t.Fail()
	}
	fmt.Println("XXXX"+fndPrdct.Uuid)
	fmt.Println(fndPrdct)
	fmt.Println("XXXX")

}
