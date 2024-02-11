package main

import (
	// "fmt"
	"fmt"
	"store/pkg/infra"
	"store/pkg/models"
	"store/pkg/repo"
	"store/pkg/service"
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
	baseProduct := seeder.BuildBasicFurnitureProduct("Base custom shelf", "A shelf build to your specifications")

	productService := service.NewProductService(repo.NewProductRepo(db), db)
	product, err := productService.CopyBaseProduct(baseProduct)
	if err != nil {
		panic(err)
	}

	// TODO: CartService...
	cart := models.NewCart()
	err = productService.AddProductToCart(cart, product)
	if err != nil {
		panic(err)
	}

	var fndCart models.Cart
	result := db.Preload("CartItems").Preload("CartItems.Product").Where("uuid = ?", cart.Uuid).First(&fndCart)

	if result.Error != nil {
		panic(result.Error)
	}
	for _, item := range fndCart.CartItems {
		fmt.Println(item.Product.Name)
	}
	fmt.Println(fndCart)

}
