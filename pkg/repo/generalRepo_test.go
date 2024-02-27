package repo

import (
	"store/pkg/infra"
	"store/pkg/models"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGeneralRepo(t *testing.T) {

	cfg := infra.ReadConfig("../../.env")
	gorm := infra.Gorm(cfg.POSTGRES_URL)

	sqlx, err := sqlx.Open("postgres", cfg.POSTGRES_URL)
	if err != nil {
		t.Fatal("failed to connect database")
	}

	repo := NewGeneralRepo(sqlx, gorm)

	// Save model.
	cartItem := models.NewCartItem()
	err = repo.Save(cartItem)
	if err != nil {
		t.Error(err)
	}

	// Retrieve model.
	var foundCartItem models.CartItem
	err = repo.GetByPrimaryKey(&foundCartItem, cartItem.Uuid)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, cartItem.Uuid, foundCartItem.Uuid)
}
