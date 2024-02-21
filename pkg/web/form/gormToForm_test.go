package form

import (
	"fmt"
	"store/pkg/infra"
	"store/pkg/models"
	"testing"
)

func TestGormToForm(_ *testing.T) {

	cfg := infra.ReadConfig("../../../.env")
	db := infra.Gorm(cfg.POSTGRES_URL)
	seeder := models.NewSeeder(db)
	product := seeder.BuildBasicProduct("a bookshelf", "custom built bookshelf")
	fmt.Println(GormToForm(product, db))
}
