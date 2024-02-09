package main

import (
	"store/pkg/infra"
	"store/pkg/models"
)

func main() {
	cfg := infra.ReadConfig("../../.env")
	db := infra.Gorm(cfg)

	// Miglate the schema
	_ = db.AutoMigrate(&models.User{})
	_ = db.AutoMigrate(&models.Product{})
	_ = db.AutoMigrate(&models.Cart{})
	_ = db.AutoMigrate(&models.CartItem{})

}
