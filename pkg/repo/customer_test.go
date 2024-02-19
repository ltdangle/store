package repo

import (
	"store/pkg/infra"
	"store/pkg/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomerRepo(t *testing.T) {
	cfg := infra.ReadConfig("../../.env")
	db := infra.Gorm(cfg.POSTGRES_URL)
	db.Exec("DELETE FROM users")

	repo := NewCustomerRepo(db)
	customer := models.NewCustomer()
	customer.Email = "emailNew@domain.net"

	// Insert statement on first save.
	assert.Nil(t, repo.Save(customer))

	// Find customer by uuid.
	foundCustomer, err := repo.FindByUuid(customer.Uuid)
	assert.Nil(t, err)
	assert.Equal(t, customer.Email, foundCustomer.Email)
	assert.Equal(t, customer.CreatedAt.Format(time.UnixDate), foundCustomer.CreatedAt.Format(time.UnixDate))

	// Find customer by email.
	foundCustomer, _ = repo.FindByEmail(customer.Email)
	assert.Equal(t, customer.Email, foundCustomer.Email)

	// Searching with non-existent uuid returns an error.
	_, err = repo.FindByUuid("wrong_uuid")
	assert.NotNil(t, err)

	// Delete customer.
	err = repo.Delete(customer.Uuid)
	assert.Nil(t, err)

	// Deleted customer cannot be found.
	_, err = repo.FindByUuid(customer.Uuid)
	assert.NotNil(t, err)
}
