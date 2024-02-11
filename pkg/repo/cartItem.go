package repo

import (
	"store/pkg/models"

	"gorm.io/gorm"
)

type CartItemRepo struct {
	db *gorm.DB
}

func NewCartItemRepo(db *gorm.DB) *CartItemRepo {
	return &CartItemRepo{db: db}
}

func (repo *CartItemRepo) Save(cart *models.Cart) error {
	tx := repo.db.Save(cart)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *CartItemRepo) FindByUuid(uuid string) (*models.CartItem, error) {
	var cartItem models.CartItem

	result := repo.db.Where("uuid= ?", uuid).First(&cartItem)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cartItem, nil
}

func (repo *CartItemRepo) Delete(uuid string) error {
	var cartItem models.CartItem

	result := repo.db.Where("uuid= ?", uuid).Delete(&cartItem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
