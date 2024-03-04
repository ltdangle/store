package web

import (
	"fmt"
	"net/http"
	"store/pkg/logger"
	"time"

	"github.com/gorilla/mux"
)

type ParsedUrl struct {
	Value string
	Error error
}
type AppRouter struct {
	Router *mux.Router
	Logger logger.LoggerInterface
}

func NewAppRouter(router *mux.Router, logger logger.LoggerInterface) *AppRouter {
	return &AppRouter{Router: router, Logger: logger}
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
		appRouter.Logger.Warn(fmt.Sprintf("UrlInternal: url for route %s not found ", routeName))
		return ParsedUrl{Error: fmt.Errorf("UrlInternal: url for route %s not found", routeName)}
	}

	url, err := route.URL(pairs...)
	if err != nil || url == nil {
		appRouter.Logger.Warn(fmt.Sprintf("UrlInternal: url params %s for route %s could not be parsed", pairs, routeName))
		return ParsedUrl{Error: fmt.Errorf("UrlInternal: url params %s for route %s could not be parsed", pairs, routeName)}
	}

	return ParsedUrl{Value: url.String()}
}
