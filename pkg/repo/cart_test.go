package repo

import (
	"store/pkg/infra"
	"store/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCartByUuid(t *testing.T) {
	cfg := infra.ReadConfig("../../.env")
	db := infra.Gorm(cfg.POSTGRES_URL)
	repo := NewCartRepo(db)
	cart := models.NewCart()
	cartItem := models.NewCartItem()
	cart.CartItems = append(cart.CartItems, cartItem)
	assert.Nil(t, repo.Save(cart))

	foundCart, err := repo.FindByCartItemUuid(cartItem.Uuid)
	assert.Nil(t, err)
	assert.Equal(t, cart.Uuid, foundCart.Uuid)
}
