package repo

import (
	"store/pkg/models"

	"gorm.io/gorm"
)

type CartRepo struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) *CartRepo {
	return &CartRepo{db: db}
}

func (repo *CartRepo) Save(cart *models.Cart) error {
	tx := repo.db.Save(cart)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *CartRepo) FindByUuid(uuid string) (*models.Cart, error) {
	var cart models.Cart

	result := repo.db.Where("uuid= ?", uuid).First(&cart)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cart, nil
}

func (repo *CartRepo) Delete(uuid string) error {
	var cart models.User

	result := repo.db.Where("uuid= ?", uuid).Delete(&cart)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
