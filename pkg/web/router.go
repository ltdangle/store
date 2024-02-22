package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ParsedUrl struct {
	Value string
	Error error
}
type AppRouter struct {
	Router *mux.Router
}

func NewAppRouter(router *mux.Router) *AppRouter {
	return &AppRouter{Router: router}
}

func (appRouter *AppRouter) Response(w http.ResponseWriter, html string) {
	// Set headers to disable caching
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max- age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", time.Now().Format(http.TimeFormat))

	fmt.Fprint(w, html)

}
func (appRouter *AppRouter) UrlInternal(routeName string, pairs ...string) ParsedUrl {
	route := appRouter.Router.Get(routeName)
	if route == nil {
		return ParsedUrl{Error: fmt.Errorf("UrlInternal: url for route %s not found", routeName)}
	}

	url, err := route.URL(pairs...)
	if err != nil || url == nil {
		return ParsedUrl{Error: fmt.Errorf("UrlInternal: url params %s for route %s could not be parsed", pairs, routeName)}
	}

	return ParsedUrl{Value: url.String()}
}
