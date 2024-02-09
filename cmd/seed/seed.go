package main

import (
	"store/pkg/infra"
	"store/pkg/models"
)

func main() {
	cfg := infra.ReadConfig(".env")
	db := infra.Gorm(cfg)
	// Truncate tables.
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM products")
	db.Exec("DELETE FROM carts")
	db.Exec("DELETE FROM cart_items")

	// Init seeder.
	seeder := models.NewSeeder(db)
	customer := seeder.CreateCustomer()
	cart := seeder.AddCart(customer)
	cartItem := seeder.AddCartItem(cart)
	product := seeder.CreateProduct()
	seeder.AddProduct(cartItem, product)

}
