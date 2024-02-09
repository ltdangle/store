package service

import (
	"store/pkg/models"
	"store/pkg/repo"
)

type CustomerService struct {
	repo *repo.CustomerRepo
}

func NewCustomerService(repo *repo.CustomerRepo) *CustomerService {
	return &CustomerService{repo: repo}
}

type CreateCustomerRqst struct {
	Email string // TODO: validate on empty, on length
}

func (service *CustomerService) Create(rqst CreateCustomerRqst) (*models.User, error) {
	customer := models.NewCustomer()
	customer.Email = rqst.Email

	err := service.repo.Save(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
