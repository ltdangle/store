package repo

import (
	"store/pkg/models"

	"gorm.io/gorm"
)

type BaseProductRepo struct {
	db *gorm.DB
}

func NewBaseProductRepo(db *gorm.DB) *BaseProductRepo {
	return &BaseProductRepo{db: db}
}

func (repo *BaseProductRepo) Save(product *models.BaseProduct) error {
	tx := repo.db.Save(product)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *BaseProductRepo) FindByUuid(uuid string) (*models.BaseProduct, error) {
	var product models.BaseProduct

	result := repo.db.Preload("Fields").Where("uuid= ?", uuid).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (repo *BaseProductRepo) Delete(uuid string) error {
	var product models.BaseProduct

	result := repo.db.Where("uuid= ?", uuid).Delete(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
