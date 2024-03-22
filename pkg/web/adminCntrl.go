package web

import (
	"fmt"
	"net/http"
	"store/pkg/i"
	"store/pkg/logger"
	"store/pkg/repo"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type EntityToTableMap struct {
	Entity    any
	TableName string
}

type AdminController struct {
	router *AppRouter
	logger logger.LoggerInterface
	repo   *repo.GeneralRepo

	// map[EntityName]EntityObject
	mappedEntities map[string]i.AdminEntity
}

func NewAdminController(router *AppRouter, logger logger.LoggerInterface, r *repo.GeneralRepo) *AdminController {
	return &AdminController{router: router, logger: logger, repo: r, mappedEntities: make(map[string]i.AdminEntity)}
}

func (cntrl *AdminController) AddMappedEntity(key string, entity i.AdminEntity) {
	cntrl.mappedEntities[key] = entity
}

// Views all mapped entity.
const ADMIN_VIEW_ALL_ENTITIES_ROUTE = "admin view all entities route"

func (cntrl *AdminController) ViewAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityName := vars["entity"]

	entity, ok := cntrl.mappedEntities[entityName]
	if !ok {
		cntrl.router.RndrTmpl(w, "Entity type not found")
		return
	}

	query := fmt.Sprintf(`SELECT * FROM %s;`, entity.TableName())
	resultsMap, err := cntrl.repo.QueryToMap(query)
	if err != nil {
		cntrl.router.RndrTmpl(w, err.Error())
		return
	}

	table := NewAdminTable(entityName)
	table.DataMap = resultsMap
	cntrl.router.RndrTmpl(w, table.Render(cntrl.router))
}

// Views mapped entity.
const ADMIN_VIEW_ENTITY_ROUTE = "admin view entity route"

func (cntrl *AdminController) ViewEntity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityName := vars["entity"]
	uuid := vars["uuid"]

	mappedEntity, ok := cntrl.mappedEntities[entityName]
	if !ok {
		cntrl.router.RndrTmpl(w, "Entity type not found")
		return
	}

	// Retrieve mapped entity.
	err := cntrl.repo.GetByPrimaryKey(mappedEntity, uuid)
	if err != nil {
		cntrl.router.RndrTmpl(w, "Entity type not found")
		return
	}

	// Populate form.
	form, err := AdminForm(mappedEntity)
	if err != nil {
		cntrl.router.Response(w, err.Error())
		return
	}

	form.Action = cntrl.router.UrlInternal(ADMIN_UPDATE_ENTITY_ROUTE, "entity", entityName, "uuid", uuid)
	cntrl.router.RndrTmpl(w, form.Render())
}

const ADMIN_UPDATE_ENTITY_ROUTE = "admin update entity route"

// Updates mapped entity.
func (cntrl *AdminController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityName := vars["entity"]
	uuid := vars["uuid"]

	entityPointer, ok := cntrl.mappedEntities[entityName]
	if !ok {
		cntrl.router.RndrTmpl(w, "Entity type not found")
		return
	}

	if err := r.ParseForm(); err != nil {
		cntrl.router.Response(w, err.Error())
		return
	}

	// Convert request to struct using gorilla schema.
	var decoder = schema.NewDecoder()
	err := decoder.Decode(entityPointer, r.PostForm)
	if err != nil {
		cntrl.logger.Warn(err)
		cntrl.router.Response(w, err.Error())
		return
	}

	// Save entity.
	err = cntrl.repo.Save(entityPointer)
	if err != nil {
		cntrl.router.Response(w, err.Error())
		cntrl.logger.Warn(err)
		return
	}

	redirectUrl := cntrl.router.UrlInternal(ADMIN_VIEW_ENTITY_ROUTE, "entity", entityName, "uuid", uuid)
	http.Redirect(w, r, redirectUrl, http.StatusFound)
}

const ADMIN_CREATE_ENTITY_ROUTE = "admin create entity route"

// Updates mapped entity.
func (cntrl *AdminController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityName := vars["entity"]

	mappedEntity, ok := cntrl.mappedEntities[entityName]
	if !ok {
		cntrl.router.RndrTmpl(w, "Entity type not found")
		return
	}

	// Populate form.
	form, err := AdminForm(mappedEntity)
	if err != nil {
		cntrl.router.Response(w, err.Error())
		return
	}

	form.Action = cntrl.router.UrlInternal(ADMIN_UPDATE_ENTITY_ROUTE, "entity", entityName, "uuid", "xxx")
	cntrl.router.RndrTmpl(w, form.Render())
}
