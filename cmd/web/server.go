package main

import (
	"net/http"
	"store/pkg/dc"

	log "github.com/sirupsen/logrus"
)

func main() {
	dc := dc.NewDc(".env")

	dc.Router.HandleFunc("/cart/{uuid}", dc.CartController.View).Methods("GET").Name("cart")

	log.Info("Starting server on localhost:8080")
	err := http.ListenAndServe("localhost:8080", dc.Router)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
