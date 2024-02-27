package repo

import (
	"store/pkg/models"

	"gorm.io/gorm"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(db *gorm.DB) *CustomerRepo {
	return &CustomerRepo{db: db}
}
func (repo *CustomerRepo) Save(customer *models.User) error {
	tx := repo.db.Save(customer)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}
func (repo *CustomerRepo) FindByUuid(uuid string) (*models.User, error) {
	var customer models.User

	result := repo.db.Where("uuid= ? AND type= ?", uuid, models.USER_CUSTOMER).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (repo *CustomerRepo) FindByEmail(email string) (*models.User, error) {
	var customer models.User

	result := repo.db.Where("email= ? AND type= ?", email, models.USER_CUSTOMER).First(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (repo *CustomerRepo) Delete(uuid string) error {
	var customer models.User

	result := repo.db.Where("uuid= ?", uuid).Delete(&customer)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
