package main

import (
	"store/pkg/infra"
	"store/pkg/models"
	"store/pkg/repo"
	"store/pkg/service"

	"github.com/bxcodec/faker/v3"
)

func main() {
	cfg := infra.ReadConfig(".env")
	db := infra.Gorm(cfg)
	// Truncate tables.
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM products")
	db.Exec("DELETE FROM carts")
	db.Exec("DELETE FROM cart_items")

	// Int customer service.
	cstmrRpo := repo.NewCustomerRepo(db)
	cstmrSrvc := service.NewCustomerService(cstmrRpo)
	customer, _ := cstmrSrvc.Create(service.CreateCustomerRqst{Email: faker.Email()})

	// Init seeder.
	seeder := models.NewSeeder(db)
	cart := seeder.AddCart(customer)
	cartItem := seeder.AddCartItem(cart)
	product := seeder.CreateProduct()
	seeder.AddProduct(cartItem, product)

}
