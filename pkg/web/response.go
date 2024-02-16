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

func UrlInternal(router *mux.Router, routeName string, pairs ...string) (string, error) {
	route := router.Get(CART_ROUTE)
	if route == nil {
		return "", fmt.Errorf("UrlInternal: url for route %s not found", routeName)
	}

	urlStr, err := route.URL(pairs...)
	if err != nil || urlStr == nil {
		return "", fmt.Errorf("UrlInternal: url params %s for route %s could not be parsed", pairs, routeName)
	}

	return urlStr.String(), nil
}
