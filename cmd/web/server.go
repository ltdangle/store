package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"store/pkg/dc"
	"store/pkg/models"
	"store/pkg/web"

	log "github.com/sirupsen/logrus"
)

func main() {
	dc := dc.NewDc(".env")
	router := dc.AppRouter.Router
	router.Use(loggingMiddleware)

	// Admin panel.
	router.HandleFunc("/admin/{entity}/all", dc.AdminController.ViewAll).Methods("GET").Name(web.ADMIN_VIEW_ALL_ENTITIES_ROUTE)
	router.HandleFunc("/admin/{entity}/{uuid}/view", dc.AdminController.ViewEntity).Methods("GET").Name(web.ADMIN_VIEW_ENTITY_ROUTE)
	router.HandleFunc("/admin/{entity}/{uuid}/update", dc.AdminController.Update).Methods("POST").Name(web.ADMIN_UPDATE_ENTITY_ROUTE)
	router.HandleFunc("/admin/{entity}/create", dc.AdminController.Create).Methods("GET").Name(web.ADMIN_CREATE_ENTITY_ROUTE)

	dc.AdminController.AddMappedEntity("cart", models.Cart{})
	dc.AdminController.AddMappedEntity("cartItem", models.CartItem{})
	dc.AdminController.AddMappedEntity("product", models.Product{})

	dc.AdminTemplate.AddNavLink(dc.AppRouter.UrlInternal(web.ADMIN_VIEW_ALL_ENTITIES_ROUTE, "entity", "cart"), "Carts")
	dc.AdminTemplate.AddNavLink(dc.AppRouter.UrlInternal(web.ADMIN_VIEW_ALL_ENTITIES_ROUTE, "entity", "cartItem"), "Cart Items")
	dc.AdminTemplate.AddNavLink(dc.AppRouter.UrlInternal(web.ADMIN_VIEW_ALL_ENTITIES_ROUTE, "entity", "product"), "Products")

	// Store.
	router.HandleFunc("/cart/{uuid}", dc.CartController.View).Methods("GET").Name(web.CART_VIEW_ROUTE)
	router.HandleFunc("/cartItem/{uuid}/delete", dc.CartController.DeleteItem).Methods("GET").Name(web.CART_ITEM_DELETE_ROUTE)
	router.HandleFunc("/seed", seed).Methods("GET")

	loggedRouter := logAllResponsesMiddleware(router)

	log.Info("Starting server on localhost:8080")
	err := http.ListenAndServe("localhost:8080", loggedRouter)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func seed(w http.ResponseWriter, _ *http.Request) {
	// The command you want to run
	cmd := exec.Command("go", "run", "cmd/seed/seed.go")

	// Create a buffer to capture the output
	var out bytes.Buffer
	cmd.Stdout = &out

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	// Output the result
	fmt.Fprint(w, out.String())
}

func logAllResponsesMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s %s ", r.Method, r.RequestURI)
		rec := newStatusRecorder(w)
		handler.ServeHTTP(rec, r)

		// Now you can log the response status
		log.Printf("response:  %d", rec.status)
	})
}

func newStatusRecorder(w http.ResponseWriter) *responseRecorder {
	return &responseRecorder{w, http.StatusOK}
}

type responseRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *responseRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
