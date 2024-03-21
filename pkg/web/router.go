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
type Template interface {
	Render(content string) string
}
type AppRouter struct {
	Router   *mux.Router
	Logger   logger.LoggerInterface
	Template Template
}

func NewAppRouter(router *mux.Router, template Template, logger logger.LoggerInterface) *AppRouter {
	return &AppRouter{Router: router, Template: template, Logger: logger}
}

func (appRouter *AppRouter) RndrTmpl(w http.ResponseWriter, html string) {
	appRouter.Response(w, appRouter.Template.Render(html))
}

func (appRouter *AppRouter) Response(w http.ResponseWriter, html string) {
	// Set headers to disable caching
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max- age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", time.Now().Format(http.TimeFormat))

	fmt.Fprint(w, html)

}
func (appRouter *AppRouter) UrlInternal(routeName string, pairs ...string) string {
	route := appRouter.Router.Get(routeName)
	if route == nil {
		appRouter.Logger.Warn(fmt.Sprintf("UrlInternal: url for route %s not found ", routeName))
		return ""
	}

	url, err := route.URL(pairs...)
	if err != nil || url == nil {
		appRouter.Logger.Warn(fmt.Sprintf("UrlInternal: url params %s for route %s could not be parsed", pairs, routeName))
		return ""
	}

	return url.String()
}

func (appRouter *AppRouter) Render(content string) {

}
