package web

import (
	"store/pkg/logger"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AdminController struct {
	router *mux.Router
	logger logger.LoggerInterface
	tmpl   *Tmpl
	db     *gorm.DB
}

func NewAdminController(router *mux.Router, logger logger.LoggerInterface, tmpl *Tmpl, db *gorm.DB) *AdminController {
	return &AdminController{router: router, logger: logger, tmpl: tmpl, db: db}
}
