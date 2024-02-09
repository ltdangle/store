package repo

import (
	"store/pkg/infra"
	"store/pkg/models"
	"testing"
)

func TestCustomerRepo(t *testing.T) {
	cfg := infra.ReadConfig("../../.env")
	gorm := infra.Gorm(cfg)

	repo := NewCustomerRepo(gorm)
	customer := models.NewCustomer()
	repo.Save(customer)
	t.Log("test ran")
}
