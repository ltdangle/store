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
	seeder.BuildProduct("Custom table", "A table build to your specifications")
	seeder.BuildProduct("Custom shelf", "A shelf build to your specifications")

	// Int customer service.
	// cstmrRpo := repo.NewCustomerRepo(db)
	// cstmrSrvc := service.NewCustomerService(cstmrRpo)
	// customer, _ := cstmrSrvc.Create(service.CreateCustomerRqst{Email: faker.Email()})
	// fmt.Println(customer)

	// TODO: create product
	// TODO: create cart
	// TODO: create cart item

}
