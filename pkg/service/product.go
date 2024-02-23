package service

import (
	"store/pkg/models"
	"store/pkg/repo"

	"gorm.io/gorm"
)

type ProductService struct {
	repo *repo.ProductRepo
	db   *gorm.DB
}

func NewProductService(repo *repo.ProductRepo, db *gorm.DB) *ProductService {
	return &ProductService{repo: repo, db: db}
}

type NewProductRqst struct {
	Name        string
	Description string
	BasePrice   int
}

func (service *ProductService) Create(rqst NewProductRqst) (*models.Product, error) {
	product := models.NewProduct()
	product.Name = rqst.Name
	product.BasePrice = rqst.BasePrice

	err := service.repo.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductService) Save(product *models.Product) error {
	err := service.repo.Save(product)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductService) FindByUuuid(uuid string) (*models.Product, error) {
	product, err := service.repo.FindByUuid(uuid)
	if err != nil {
		return nil, err
	}

	return product, nil
}


