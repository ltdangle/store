package main

import (
	"net/http"
	"store/pkg/dc"
	"store/pkg/web"

	log "github.com/sirupsen/logrus"
)

const (
	CART_ROUTE = "cart"
)

func main() {
	dc := dc.NewDc(".env")

	dc.Router.HandleFunc("/cart/{uuid}", dc.CartController.View).Methods("GET").Name(web.CART_ROUTE)

	log.Info("Starting server on localhost:8080")
	err := http.ListenAndServe("localhost:8080", dc.Router)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
