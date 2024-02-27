package web

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"store/pkg/infra"
	"store/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type EntityToTableMap struct {
	Entity    any
	TableName string
}

type AdminController struct {
	router *AppRouter
	logger logger.LoggerInterface
	tmpl   *Tmpl
	db     *gorm.DB
	// map[EntityName]EntityObject
	mappedEntities map[string]EntityToTableMap
}

func NewAdminController(router *AppRouter, logger logger.LoggerInterface, tmpl *Tmpl, db *gorm.DB) *AdminController {
	return &AdminController{router: router, logger: logger, tmpl: tmpl, db: db,
		mappedEntities: make(map[string]EntityToTableMap),
	}
}

func (cntrl *AdminController) AddMappedEntity(key string, entity any, tableName string) {
	cntrl.mappedEntities[key] = EntityToTableMap{Entity: entity, TableName: tableName}
}

// Views mapped entity.
const ADMIN_VIEW_ENTITY_ROUTE = "admin view entity route"

func (cntrl *AdminController) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityName := vars["entity"]
	uuid := vars["uuid"]

	mappedEntity, ok := cntrl.mappedEntities[entityName]
	if !ok {
		http.Error(w, "Entity type not found", http.StatusNotFound)
		return
	}

	cfg := infra.ReadConfig(".env")
	fmt.Println(cfg)
	db, err := sqlx.Open("postgres", cfg.POSTGRES_URL)
	if err != nil {
		log.Fatal("failed to connect database")
	}

	query := fmt.Sprintf(`SELECT * FROM %s WHERE uuid = $1;`, mappedEntity.TableName)
	err = db.Get(mappedEntity.Entity, query, uuid)
	if err != nil {
		log.Fatal(err)
	}

	entityValue := reflect.ValueOf(mappedEntity.Entity).Elem().Interface()

	f := GormAdminForm(entityValue, cntrl.router)
	f.Action = cntrl.router.UrlInternal(ADMIN_UPDATE_ENTITY_ROUTE, "entity", entityName, "uuid", uuid).Value
	cntrl.tmpl.SetMain(f.Render())
	cntrl.router.Response(w, cntrl.tmpl.Render())

}

const ADMIN_UPDATE_ENTITY_ROUTE = "admin update entity route"

// Updates mapped entity.
func (cntrl *AdminController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityName := vars["entity"]
	uuid := vars["uuid"]

	entityPointer, ok := cntrl.mappedEntities[entityName]
	if !ok {
		http.Error(w, "Entity type not found", http.StatusNotFound)
		return
	}

	// entityValue := reflect.ValueOf(entityPointer).Elem().Interface()

	if err := r.ParseForm(); err != nil {
		cntrl.tmpl.SetMain(err.Error())
		cntrl.router.Response(w, cntrl.tmpl.Render())
		return
	}

	// gorilla schema
	var decoder = schema.NewDecoder()
	err := decoder.Decode(entityPointer, r.PostForm)
	if err != nil {
		cntrl.logger.Warn(err)
		cntrl.tmpl.SetMain(err.Error())
		cntrl.router.Response(w, cntrl.tmpl.Render())
		return
	}

	result := cntrl.db.Save(entityPointer)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusNotFound)
		cntrl.logger.Warn(result.Error)
		return
	}

	redirectUrl := cntrl.router.UrlInternal(ADMIN_VIEW_ENTITY_ROUTE, "entity", entityName, "uuid", uuid)
	http.Redirect(w, r, redirectUrl.Value, http.StatusFound)
}
