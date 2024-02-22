package web

import (
	"fmt"
	"net/http"
	"store/pkg/logger"
	"store/pkg/models"
	"store/pkg/repo"
	"store/pkg/service"
	"store/pkg/web/form"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"gorm.io/gorm"
)

type CartController struct {
	router  *AppRouter
	service *service.CartService
	repo    *repo.CartRepo
	logger  logger.LoggerInterface
	tmpl    *Tmpl
	db      *gorm.DB
}

func NewCartController(router *AppRouter, service *service.CartService, repo *repo.CartRepo, logger logger.LoggerInterface, tmpl *Tmpl, db *gorm.DB) *CartController {
	return &CartController{router: router, service: service, repo: repo, logger: logger, tmpl: tmpl, db: db}
}

type CartVM struct {
	Cart *models.Cart
}

const CART_VIEW_ROUTE = "cart"

func (cntrl *CartController) View(w http.ResponseWriter, r *http.Request) {

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
		cartTmpl := NewCartTmpl(cntrl.router)
		cntrl.tmpl.SetMain(cartTmpl.cart(vm))
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
	if cartUrl.Error != nil {
		cntrl.logger.Warn("CartController.DeleteItem: could not parse redirect url")
	}

	cntrl.logger.Info(fmt.Sprintf("CartController.DeleteItem: item deleted, redirect to %s", cartUrl))

	http.Redirect(w, r, cartUrl.Value, http.StatusTemporaryRedirect)
}

const CART_EDIT_ROUTE = "edit cart"

func (cntrl *CartController) EditCart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			cntrl.tmpl.SetMain(err.Error())
			cntrl.router.Response(w, cntrl.tmpl.Render())
			return
		}

		var cart models.Cart

		// gorilla schema
		var decoder = schema.NewDecoder()
		err := decoder.Decode(&cart, r.PostForm)
		if err != nil {
			cntrl.logger.Warn(err)
			cntrl.tmpl.SetMain(err.Error())
			cntrl.router.Response(w, cntrl.tmpl.Render())
			return
		}

		err = cntrl.repo.Save(&cart)
		if err != nil {
			cntrl.logger.Warn(err)
			cntrl.tmpl.SetMain(err.Error())
			cntrl.router.Response(w, cntrl.tmpl.Render())
			return
		}

		http.Redirect(w, r, r.URL.String(), http.StatusFound)

	case http.MethodGet:
		vars := mux.Vars(r)
		uuid := vars["uuid"]

		cart, err := cntrl.repo.FindByUuid(uuid)
		if err != nil {
			cntrl.logger.Warn(fmt.Sprintf("CartController.View: cart %s : %s", uuid, err.Error()))
			fmt.Fprint(w, err.Error())
		}

		f := form.GormToForm(cart, cntrl.db)
		f.Method = "POST"
		cntrl.tmpl.SetMain(f.Render())
		cntrl.router.Response(w, cntrl.tmpl.Render())
	}
}
