package repo

import (
	"store/pkg/models"

	"gorm.io/gorm"
)

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *customerRepo {
	return &customerRepo{db: db}
}
func (repo *customerRepo) Save(customer *models.User) error {
	tx := repo.db.Save(customer)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (repo *customerRepo) FindByUuid(uuid string) (*models.User, error) {
	var customer models.User

	result := repo.db.Where("uuid= ? AND type= ?", uuid, models.USER_CUSTOMER).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (repo *customerRepo) Delete(uuid string) error {
	var customer models.User

	result := repo.db.Where("uuid= ?", uuid).Delete(&customer)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
