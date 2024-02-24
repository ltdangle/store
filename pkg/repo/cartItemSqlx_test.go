package repo

import (
	"fmt"
	"log"
	"store/pkg/infra"
	"store/pkg/models"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func TestCartItemSqlxz(t *testing.T) {
	cfg := infra.ReadConfig("../../.env")
	db, err := sqlx.Open("postgres", cfg.POSTGRES_URL)
	if err != nil {
		log.Fatal("failed to connect database")
	}
	var cartItem models.CartItem
	uuid := "klf"

	query := `
SELECT 
    uuid, cart_uuid, product_uuid, quantity, subtotal, created_at, updated_at, deleted_at  
FROM 
    cart_items
WHERE 
  cart_items.uuid = $1;`
	err = db.Get(&cartItem, query, uuid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", cartItem)
}
