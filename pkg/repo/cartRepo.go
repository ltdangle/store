package repo

import (
	"store/pkg/models"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type CartRepo struct {
	db   *gorm.DB
	sqlx *sqlx.DB
}

func NewCartRepo(db *gorm.DB, sqlx *sqlx.DB) *CartRepo {
	return &CartRepo{db: db, sqlx: sqlx}
}

func (repo *CartRepo) Save(cart *models.Cart) error {
	tx := repo.db.Save(cart)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

type CartItemVM struct {
	CartItem *models.CartItem
	Product  *models.Product
}

type CartVM struct {
	Cart      *models.Cart
	CartItems []CartItemVM
}

func (repo *CartRepo) FullCartNew(cartUuid string) (*CartVM, error) {
	cartVM := &CartVM{}
	// Retrieve cart.
	var cart models.Cart
	err := repo.sqlx.Get(&cart, `SELECT * FROM carts WHERE uuid = $1;`, cartUuid)
	if err != nil {
		return nil, err
	}

	cartVM.Cart = &cart

	// Retrieve cart items.
	var cartItems []*models.CartItem
	err = repo.sqlx.Select(&cartItems, `SELECT * FROM cart_items WHERE cart_uuid = $1;`, cartUuid)
	if err != nil {
		return nil, err
	}

	// Retrieve products for each cart item.
	for _, cartItem := range cartItems {
		var product models.Product
		err = repo.sqlx.Get(&product, `SELECT * FROM products WHERE uuid = $1;`, cartItem.ProductUuid)
		if err != nil {
			return nil, err
		}
		cartItemVm := CartItemVM{CartItem: cartItem, Product: &product}
		cartVM.CartItems = append(cartVM.CartItems, cartItemVm)
	}
	return cartVM, nil
}

func (repo *CartRepo) FullCart(uuid string) (*models.Cart, error) {
	var cart models.Cart
	result := repo.db.Preload("CartItems").Preload("CartItems.Product").Preload("CartItems.Product.Fields").Where("uuid = ?", uuid).First(&cart)
	if result.Error != nil {
		return nil, result.Error
	}

	return &cart, nil
}
func (repo *CartRepo) FindByCartItemUuid(cartItemUuid string) (*models.Cart, error) {
	var cart models.Cart

	err := repo.db.Joins("JOIN cart_items ON cart_items.cart_id = carts.id").
		Where("cart_items.uuid = ?", cartItemUuid).
		First(&cart).Error
	if err != nil {
		return nil, err
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
