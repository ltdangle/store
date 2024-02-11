package service

import (
	"store/pkg/models"
	"store/pkg/repo"
)

type ProductService struct {
	repo *repo.ProductRepo
}

func NewProductService(repo *repo.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

type NewProductRqst struct {
	Name        string
	Description string
	Price       int
}

func (service *ProductService) Create(rqst NewProductRqst) (*models.Product, error) {
	product := models.NewProduct()
	product.Name = rqst.Name
	product.Price = rqst.Price

	err := service.repo.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductService) FindByUuuid(uuid string) (*models.Product, error) {
	product, err := service.repo.FindByUuid(uuid)
	if err != nil {
		return nil, err
	}

	return product, nil
}
