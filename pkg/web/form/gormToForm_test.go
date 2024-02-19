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
	GormToForm(models.NewCart(), db)
	fmt.Println(db)
}
