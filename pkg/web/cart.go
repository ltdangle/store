package web

import (
	"context"
	"fmt"
	"net/http"
	"store/pkg/repo"
	"store/pkg/service"
	"store/pkg/web/tmpl"

	"github.com/gorilla/mux"
)

type CartController struct {
	service *service.CartService
	repo    *repo.CartRepo
}

func NewCartController(service *service.CartService, repo *repo.CartRepo) *CartController {
	return &CartController{service: service, repo: repo}
}
func (cntrl *CartController) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	// TODO: Validate uuid.
	cart, err := cntrl.repo.FindByUuid(uuid)

	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		component := tmpl.Hello(cart)
		_ = component.Render(context.Background(), w)
	}

	// w.WriteHeader(http.StatusOK)

}
