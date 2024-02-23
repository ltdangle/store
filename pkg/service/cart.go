package service

import (
	"errors"
	"store/pkg/models"
	"store/pkg/repo"

	"gorm.io/gorm"
)

type CartService struct {
	repo         *repo.CartRepo
	cartItemRepo *repo.CartItemRepo
	db           *gorm.DB
}

func NewCartService(repo *repo.CartRepo, cartItemRepo *repo.CartItemRepo, db *gorm.DB) *CartService {
	return &CartService{repo: repo, cartItemRepo: cartItemRepo, db: db}
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
	cartItem.CartUuid = cart.Uuid
	cartItem.ProductUuid = product.Uuid

	tx := service.db.Save(cart)
	if tx.Error != nil {
		return tx.Error
	}

	tx = service.db.Save(cartItem)
	if tx.Error != nil {
		return tx.Error
	}

	tx = service.db.Save(product)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (service *CartService) RemoveCartItem(cartItemUuid string) (*models.Cart, error) {
	// Make sure active cart exists.
	cart, err := service.repo.FindByCartItemUuid(cartItemUuid)
	if err != nil {
		return nil, errors.New("CartService: cart with cartItem " + cartItemUuid + " not found")
	}

	// Delete cart item.
	// TODO: check if it belongs to a user
	err = service.cartItemRepo.Delete(cartItemUuid)
	if err != nil {
		return nil, errors.New("CartService: cartItem " + cartItemUuid + "could not be deleted")
	}
	return cart, nil
}
