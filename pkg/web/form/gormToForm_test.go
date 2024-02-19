package form

import (
	"fmt"
	"store/pkg/infra"
	"store/pkg/models"
	"testing"
)

func TestGormToForm(t *testing.T) {

	cfg := infra.ReadConfig("../../../.env")
	db := infra.Gorm(cfg.POSTGRES_URL)
	cart := models.NewCart()
	cart.ID = 1
	cart.Total = 234
	fmt.Println(GormToForm(cart, db))
}
