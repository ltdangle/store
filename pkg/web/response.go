package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func response(w http.ResponseWriter, html string) {
	// Set headers to disable caching
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max- age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", time.Now().Format(http.TimeFormat))

	fmt.Fprint(w, html)

}

type Url struct {
	Value string
	Error error
}

func UrlInternal(router *mux.Router, routeName string, pairs ...string) Url {
	route := router.Get(routeName)
	if route == nil {
		return Url{Error: fmt.Errorf("UrlInternal: url for route %s not found", routeName)}
	}

	url, err := route.URL(pairs...)
	if err != nil || url == nil {
		return Url{Error: fmt.Errorf("UrlInternal: url params %s for route %s could not be parsed", pairs, routeName)}
	}

	return Url{Value: url.String()}
}
