package main

import (
	"store/pkg/infra"
	"store/pkg/models"
	"store/pkg/repo"
	"store/pkg/service"

	"gorm.io/gorm"
	// "github.com/bxcodec/faker/v3"
)

func main() {
	cfg := infra.ReadConfig(".env")
	db := infra.Gorm(cfg.POSTGRES_URL)

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
	baseProducts := []*models.BaseProduct{
		seeder.BuildBasicProduct("shelf", "A shelf build to your specifications"),
		seeder.BuildBasicProduct("chair", "A chair build to your specifications"),
		seeder.BuildBasicProduct("table", "A table build to your specifications"),
		seeder.BuildBasicProduct("sofa", "A sofa build to your specifications"),
	}

	productService := service.NewProductService(repo.NewProductRepo(db), db)
	var products []*models.Product
	for _, baseProduct := range baseProducts {
		product, err := productService.CopyBaseProduct(baseProduct)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	cartRepo := repo.NewCartRepo(db)
	cartItemRepo := repo.NewCartItemRepo(db)
	cartService := service.NewCartService(cartRepo, cartItemRepo)
	cart, err := cartService.CreateCart()
	if err != nil {
		panic(err)
	}
	// add products to cart
	for _, product := range products {
		err = cartService.AddProductToCart(cart, product)
		if err != nil {
			panic(err)
		}
	}
	// set cart item subtotal
	// for _, item := range cart.CartItems {
	// 	item.Subtotal = 35
	// }

	tx := db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cart)
	if tx.Error != nil {
		panic(err)
	}
}
