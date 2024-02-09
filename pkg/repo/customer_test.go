package repo

import (
	"store/pkg/infra"
	"store/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerRepo(t *testing.T) {
	cfg := infra.ReadConfig("../../.env")
	db := infra.Gorm(cfg)
	db.Exec("DELETE FROM users")

	repo := NewCustomerRepo(db)
	customer := models.NewCustomer()
	customer.Email = "emailNew@domain.net"

	// Insert statement on first save.
	assert.Nil(t, repo.Save(customer))

	// Saving new customer with the same email triggers an error.
	duplicateCustomer := models.NewCustomer()
	duplicateCustomer.Email = "emailNew@domain.net"
	assert.NotNil(t, repo.Save(duplicateCustomer))

	// Find customer.
	foundCustomer, err := repo.FindByUuid(customer.Uuid)
	assert.Nil(t, err)
	assert.Equal(t, customer.Email, foundCustomer.Email)
}
