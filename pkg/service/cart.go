package service

import (
	"store/pkg/models"
	"store/pkg/repo"

)

type CartService struct {
	repo *repo.CartRepo
}

func NewCartService(repo *repo.CartRepo) *CartService {
	return &CartService{repo: repo}
}

func (service *CartService) CreateCart() (*models.Cart, error) {
	cart := models.NewCart()
	err := service.repo.Save(cart)
	if err != nil {
		return nil, err
	}
	return cart, err
}

func (service *CartService) AddProductToCart(cart *models.Cart, product *models.Product) error {
	cartItem := models.NewCartItem()
	cartItem.Product = product
	cart.CartItems = append(cart.CartItems, cartItem)

	error := service.repo.Save(cart)
	if error != nil {
		return error
	}

	return nil
}
