package web

import (
	"net/http"
	"reflect"
	"store/pkg/logger"
	"store/pkg/web/form"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdminController struct {
	router *AppRouter
	logger logger.LoggerInterface
	tmpl   *Tmpl
	db     *gorm.DB
	// map[EntityName]EntityObject
	mappedEntities map[string]any
}

func NewAdminController(router *AppRouter, logger logger.LoggerInterface, tmpl *Tmpl, db *gorm.DB) *AdminController {
	return &AdminController{router: router, logger: logger, tmpl: tmpl, db: db,
		mappedEntities: map[string]any{
			// "cart":     &models.Cart{},
			// "cartItem": &models.CartItem{},
		},
	}
}

// Views mapped entity.
const ADMIN_VIEW_ENTITY_ROUTE = "admin view entity route"

func (cntrl *AdminController) AddMappedEntity(key string, entity any) {
	cntrl.mappedEntities[key] = entity
}
func (cntrl *AdminController) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityName := vars["entity"]
	uuid := vars["uuid"]

	entityPointer, ok := cntrl.mappedEntities[entityName]
	if !ok {
		http.Error(w, "Entity type not found", http.StatusNotFound)
		return
	}

	// TODO: check for result error and 0 returned results.

	result := cntrl.db.Preload(clause.Associations).Where("uuid = ?", uuid).First(entityPointer)
	if result.Error != nil {
		cntrl.logger.Warn(result.Error)
	}

	entityValue := reflect.ValueOf(entityPointer).Elem().Interface()

	f := form.GormToForm(entityValue, cntrl.db)
	f.Method = "POST"
	cntrl.tmpl.SetMain(f.Render())
	cntrl.router.Response(w, cntrl.tmpl.Render())

}

// Updates mapped entity.
func (cntrl *AdminController) Update(w http.ResponseWriter, r *http.Request) {

}
