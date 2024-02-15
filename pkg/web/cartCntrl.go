package web

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"store/pkg/models"
	"store/pkg/repo"
	"store/pkg/service"
	"time"

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
	Cart *models.Cart
}

func (cntrl *CartController) View(w http.ResponseWriter, r *http.Request) {
	// Set headers to disable caching
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max- age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", time.Now().Format(http.TimeFormat))

	var html bytes.Buffer

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	// TODO: Validate uuid.
	cart, err := cntrl.repo.FindByUuid(uuid)

	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		vm := CartVM{
			Cart: cart,
		}

		_ = Template(vm).Render(context.Background(), &html)
		fmt.Fprint(w, html.String())
	}
}
func (cntrl *CartController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemUuid := vars["uuid"]
	err := cntrl.service.RemoveCartItem(cartItemUuid)
	if err != nil {
		//TODO: log error
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Deleted")
}
