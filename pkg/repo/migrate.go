package repo

import (
	"store/pkg/infra"
	"store/pkg/models"
)

func Migrate(envFile string) {
	cfg := infra.ReadConfig(envFile)
	db := infra.Gorm(cfg.POSTGRES_URL)

	// Miglate the schema
	_ = db.AutoMigrate(&models.User{})
	_ = db.AutoMigrate(&models.BaseProduct{})
	_ = db.AutoMigrate(&models.BaseProductField{})
	_ = db.AutoMigrate(&models.Product{})
	_ = db.AutoMigrate(&models.ProductField{})
	_ = db.AutoMigrate(&models.Cart{})
	// 	db.Exec(`ALTER TABLE carts
	//     			ADD CONSTRAINT fk_user
	//     			FOREIGN KEY (user_uuid)
	//     			REFERENCES users (uuid)
	//     			ON UPDATE CASCADE;
	// `)
	_ = db.AutoMigrate(&models.CartItem{})
	// 	db.Exec(`ALTER TABLE cart_items
	//     			ADD CONSTRAINT fk_cart
	//     			FOREIGN KEY (cart_uuid)
	//     			REFERENCES carts (uuid)
	//     			ON UPDATE CASCADE;
	// `)
	// 	db.Exec(`ALTER TABLE cart_items
	//     			ADD CONSTRAINT fk_product
	//     			FOREIGN KEY (product_uuid)
	//     			REFERENCES products (uuid)
	//     			ON UPDATE CASCADE;
	// `)

}
