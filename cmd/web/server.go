package main

import (
	"fmt"
	"net/http"
	"store/pkg/dc"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	dc := dc.NewDc(".env")

	router := mux.NewRouter()

	router.HandleFunc("/cart/{uuid}", dc.CartController.View).Methods("GET")

	log.Info("Starting server on localhost:8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, uuid)
}
