package main

import (
	// "fmt"
	"store/pkg/infra"
	"store/pkg/models"
	"store/pkg/repo"
	// "store/pkg/service"
	// "github.com/bxcodec/faker/v3"
)

func main() {
	cfg := infra.ReadConfig(".env")
	db := infra.Gorm(cfg)

	// Drop tables.
	tables := []string{
		"users",
		"base_products",
		"base_product_fields",
		"products",
		"product_fields",
		"carts",
		"cart_items",
	}
	for _, tbl := range tables {
		query := "DROP TABLE IF EXISTS " + tbl + " CASCADE"
		tx := db.Exec(query)
		if tx.Error != nil {
			panic(tx.Error)
		}
	}

	// Migrate db.
	repo.Migrate(".env")

	// Seed products.
	seeder := models.NewSeeder(db)
	seeder.BuildBasicFurnitureProduct("Base custom table", "A table build to your specifications")
	seeder.BuildBasicFurnitureProduct("Base custom shelf", "A shelf build to your specifications")

	seeder.BuildFurnitureProduct("Custom table", "A table build to your specifications")
	product:=seeder.BuildFurnitureProduct("Custom shelf", "A shelf build to your specifications")

	cart := models.NewCart()
	lineItem1 := models.NewCartItem()
	lineItem1.Product = *product
	cart.CartItems = append(cart.CartItems, lineItem1)
	db.Save(product)
	db.Save(cart)
	db.Save(lineItem1)

	// Int customer service.
	// cstmrRpo := repo.NewCustomerRepo(db)
	// cstmrSrvc := service.NewCustomerService(cstmrRpo)
	// customer, _ := cstmrSrvc.Create(service.CreateCustomerRqst{Email: faker.Email()})
	// fmt.Println(customer)

	// TODO: create product
	// TODO: create cart
	// TODO: create cart item

}
