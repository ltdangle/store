package web

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"store/pkg/logger"
	"store/pkg/models"
	"store/pkg/repo"
	"store/pkg/service"

	"github.com/gorilla/mux"
)

type CartController struct {
	router  *mux.Router
	service *service.CartService
	repo    *repo.CartRepo
	logger  logger.LoggerInterface
}

func NewCartController(router *mux.Router, service *service.CartService, repo *repo.CartRepo, logger logger.LoggerInterface) *CartController {
	return &CartController{router: router, service: service, repo: repo, logger: logger}
}

type CartVM struct {
	Cart *models.Cart
}

const CART_VIEW_ROUTE = "cart"

func (cntrl *CartController) View(w http.ResponseWriter, r *http.Request) {
	var html bytes.Buffer

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	cart, err := cntrl.repo.FindByUuid(uuid)

	if err != nil {
		cntrl.logger.Warn(fmt.Sprintf("CartController.View: cart %s : %s", uuid, err.Error()))
		fmt.Fprint(w, err.Error())
	} else {
		vm := CartVM{
			Cart: cart,
		}

		_ = Template(cntrl.router, vm).Render(context.Background(), &html)
		response(w, html.String())
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

	cartUrl := UrlInternal(cntrl.router, CART_VIEW_ROUTE, "uuid", cart.Uuid)
	if cartUrl.Error != nil {
		cntrl.logger.Warn("CartController.DeleteItem: could not parse redirect url")
	}

	cntrl.logger.Info(fmt.Sprintf("CartController.DeleteItem: item deleted, redirect to %s", cartUrl))

	http.Redirect(w, r, cartUrl.Value, http.StatusTemporaryRedirect)
}
