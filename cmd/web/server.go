package main

import (
	"net/http"
	"store/pkg/dc"
	"store/pkg/web"

	log "github.com/sirupsen/logrus"
)

const (
	CART_ROUTE             = "cart"
	CART_ITEM_DELETE_ROUTE = "delete cart item"
)

func main() {
	dc := dc.NewDc(".env")

	dc.Router.HandleFunc("/cart/{uuid}", dc.CartController.View).Methods("GET").Name(web.CART_ROUTE)
	dc.Router.HandleFunc("/cartItem/{uuid}/delete", dc.CartController.DeleteItem).Methods("GET").Name(web.CART_ITEM_DELETE_ROUTE)

	log.Info("Starting server on localhost:8082")
	err := http.ListenAndServe("localhost:8082", dc.Router)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
