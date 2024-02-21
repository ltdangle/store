package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"store/pkg/dc"
	"store/pkg/web"

	log "github.com/sirupsen/logrus"
)

func main() {
	dc := dc.NewDc(".env")

	dc.Router.Use(loggingMiddleware)
	dc.Router.HandleFunc("/cart/{uuid}", dc.CartController.View).Methods("GET").Name(web.CART_VIEW_ROUTE)
	dc.Router.HandleFunc("/cartItem/{uuid}/delete", dc.CartController.DeleteItem).Methods("GET").Name(web.CART_ITEM_DELETE_ROUTE)
	dc.Router.HandleFunc("/cartItem/{uuid}/edit", dc.CartController.EditCart).Name(web.CART_EDIT_ROUTE)
	dc.Router.HandleFunc("/seed", seed).Methods("GET")

	loggedRouter := logAllResponsesMiddleware(dc.Router)

	log.Info("Starting server on localhost:8080")
	err := http.ListenAndServe("localhost:8080", loggedRouter)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func seed(w http.ResponseWriter, r *http.Request) {
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
		rec := newStatusRecorder(w)
		handler.ServeHTTP(rec, r)

		// Now you can log the response status
		log.Printf("%s %s %d", r.Method, r.RequestURI, rec.status)
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
