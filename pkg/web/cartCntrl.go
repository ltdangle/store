package web

import (
	"fmt"
	"net/http"
	"store/pkg/logger"
	"store/pkg/models"
	"store/pkg/repo"
	"store/pkg/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CartController struct {
	router  *AppRouter
	service *service.CartService
	repo    *repo.CartRepo
	logger  logger.LoggerInterface
	tmpl    *AdminTmpl
	db      *gorm.DB
}

func NewCartController(router *AppRouter, service *service.CartService, repo *repo.CartRepo, logger logger.LoggerInterface, tmpl *AdminTmpl, db *gorm.DB) *CartController {
	return &CartController{router: router, service: service, repo: repo, logger: logger, tmpl: tmpl, db: db}
}

type CartVM struct {
	Cart *models.Cart
}

const CART_VIEW_ROUTE = "cart"

func (cntrl *CartController) View(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	_, err := cntrl.repo.FindByUuid(uuid)

	if err != nil {
		cntrl.logger.Warn(fmt.Sprintf("CartController.View: cart %s : %s", uuid, err.Error()))
		fmt.Fprint(w, err.Error())
	} else {
		vm := VM{}
		cntrl.tmpl.SetMain(cntrl.tmpl.cart(vm))
		cntrl.router.Response(w, cntrl.tmpl.Render())
	}
}

const CART_ITEM_DELETE_ROUTE = "delete cart item"

func (cntrl *CartController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemUuid := vars["uuid"]

	cart, err := cntrl.service.RemoveCartItem(cartItemUuid)
	if err != nil {
		cntrl.logger.Warn(fmt.Sprintf("CartController.DeleteItem: cart with cartItem %s : %s", cartItemUuid, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cartUrl := cntrl.router.UrlInternal(CART_VIEW_ROUTE, "uuid", cart.Uuid)

	cntrl.logger.Info(fmt.Sprintf("CartController.DeleteItem: item deleted, redirect to %s", cartUrl))

	http.Redirect(w, r, cartUrl, http.StatusTemporaryRedirect)
}
