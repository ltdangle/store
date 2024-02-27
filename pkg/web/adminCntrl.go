package web

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"store/pkg/infra"
	"store/pkg/logger"
	"store/pkg/repo"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type EntityToTableMap struct {
	Entity    any
	TableName string
}

type AdminController struct {
	router *AppRouter
	logger logger.LoggerInterface
	tmpl   *Tmpl
	repo   *repo.GeneralRepo

	// map[EntityName]EntityObject
	mappedEntities map[string]EntityToTableMap
}

func NewAdminController(router *AppRouter, logger logger.LoggerInterface, tmpl *Tmpl, repo *repo.GeneralRepo) *AdminController {
	return &AdminController{router: router, logger: logger, tmpl: tmpl, repo: repo,
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

	// Convert request to struct using gorilla schema.
	var decoder = schema.NewDecoder()
	err := decoder.Decode(entityPointer.Entity, r.PostForm)
	if err != nil {
		cntrl.logger.Warn(err)
		cntrl.tmpl.SetMain(err.Error())
		cntrl.router.Response(w, cntrl.tmpl.Render())
		return
	}

	// Save entity.
	err = cntrl.repo.Save(entityPointer.Entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		cntrl.logger.Warn(err)
		return
	}

	redirectUrl := cntrl.router.UrlInternal(ADMIN_VIEW_ENTITY_ROUTE, "entity", entityName, "uuid", uuid)
	http.Redirect(w, r, redirectUrl.Value, http.StatusFound)
}
