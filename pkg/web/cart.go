package web

import (
	"fmt"
	"net/http"
	"store/pkg/repo"
	"store/pkg/service"

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

	var output string
	if err != nil {
		output = err.Error()
	} else {
		output = fmt.Sprintf("Cart %s created at %s", cart.Uuid, cart.CreatedAt.String())
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, output)
}
