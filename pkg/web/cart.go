package web

import (
	"fmt"
	"net/http"
	"store/pkg/service"

	"github.com/gorilla/mux"
)

type CartController struct {
	service *service.CartService
}

func NewCartController(service *service.CartService) *CartController {
	return &CartController{service: service}
}
func (cntrl *CartController) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, uuid)
}
