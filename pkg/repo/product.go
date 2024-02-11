package repo

import (
	"store/pkg/models"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (repo *ProductRepo) Save(product *models.Product) error {
	tx := repo.db.Save(product)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *ProductRepo) FindByUuid(uuid string) (*models.Product, error) {
	var product models.Product

	result := repo.db.Where("uuid= ?", uuid).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (repo *ProductRepo) Delete(uuid string) error {
	var product models.Product

	result := repo.db.Where("uuid= ?", uuid).Delete(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
