package web

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"store/pkg/repo"
	"store/pkg/service"
	"store/pkg/web/tmpl"

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

func (cntrl *CartController) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	// TODO: Validate uuid.
	cart, err := cntrl.repo.FindByUuid(uuid)

	var html bytes.Buffer

	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		_ = tmpl.Template(cart).Render(context.Background(), &html)

		// TODO: extract route passing logic; 1. check router.Get() for nil; 2. router.URL() for error.
		cartUrl, err := cntrl.router.Get("cart").URL("uuid", uuid)
		if err != nil {
			//TODO: log error
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println("Cart route: " + cartUrl.String())

		fmt.Fprint(w, html.String())
	}

}
