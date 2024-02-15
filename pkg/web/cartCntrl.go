package web

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"store/pkg/models"
	"store/pkg/repo"
	"store/pkg/service"

	"github.com/gorilla/mux"
)

type CartController struct {
	router  *mux.Router
	service *service.CartService
	repo    *repo.CartRepo
}

func NewCartController(router *mux.Router, service *service.CartService, repo *repo.CartRepo) *CartController {
	return &CartController{router: router, service: service, repo: repo}
}

type CartVM struct {
	Cart         *models.Cart
	DelteCartUrl string
}

func (cntrl *CartController) View(w http.ResponseWriter, r *http.Request) {
	var html bytes.Buffer

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	// TODO: Validate uuid.
	cart, err := cntrl.repo.FindByUuid(uuid)

	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {

		// TODO: extract route passing logic; 1. check router.Get() for nil; 2. router.URL() for error.
		cartUrl, err := cntrl.router.Get(CART_ITEM_DELETE_ROUTE).URL("uuid", uuid)
		if err != nil {
			//TODO: log error
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		vm := CartVM{
			Cart:         cart,
			DelteCartUrl: cartUrl.String(),
		}

		fmt.Println("Cart route: " + cartUrl.String())
		_ = Template(vm).Render(context.Background(), &html)
		fmt.Fprint(w, html.String())
	}
}
func (cntrl *CartController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemUuid:= vars["cartItemUuid"]
	err := cntrl.service.RemoveCartItem(cartItemUuid)
	if err != nil {
		//TODO: log error
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Deleted")
}
