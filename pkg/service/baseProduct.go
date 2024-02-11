package service

import (
	"store/pkg/models"
	"store/pkg/repo"
)

type BaseProductService struct {
	repo *repo.BaseProductRepo
}

func NewBaseProductService(repo *repo.BaseProductRepo) *BaseProductService {
	return &BaseProductService{repo: repo}
}

type NewBaseProductRqst struct {
	Name        string
	Description string
	BasePrice   int
}

func (service *BaseProductService) Create(rqst NewBaseProductRqst) (*models.BaseProduct, error) {
	product := models.NewBaseProduct()
	product.Name = rqst.Name
	product.BasePrice = rqst.BasePrice

	err := service.repo.Save(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *BaseProductService) Save(product *models.BaseProduct) error {
	err := service.repo.Save(product)
	if err != nil {
		return err
	}
	return nil
}

func (service *BaseProductService) FindByUuuid(uuid string) (*models.BaseProduct, error) {
	product, err := service.repo.FindByUuid(uuid)
	if err != nil {
		return nil, err
	}

	return product, nil
}
