package repo

import (
	"errors"
	"store/pkg/models"

	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *customerRepo {
	return &customerRepo{db: db}
}
func (repo *customerRepo) Save(customer *models.User) {
	repo.db.Save(customer)
}
func (repo *customerRepo) FindByUuid(uuid string) (*models.User, error) {
	var customer models.User

	result := repo.db.Where("uuid= ? AND type= ?", uuid, models.USER_CUSTOMER).First(&customer)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("customer not found")
		} else {
			return nil, result.Error
		}
	}

	return &customer, nil
}
