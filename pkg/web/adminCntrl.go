package web

import (
	"fmt"
	"net/http"
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
	tmpl   *AdminTmpl
	repo   *repo.GeneralRepo

	// map[EntityName]EntityObject
	mappedEntities map[string]repo.MappedEntity
}

func NewAdminController(router *AppRouter, logger logger.LoggerInterface, tmpl *AdminTmpl, r *repo.GeneralRepo) *AdminController {
	return &AdminController{router: router, logger: logger, tmpl: tmpl, repo: r,
		mappedEntities: make(map[string]repo.MappedEntity),
	}
}

func (cntrl *AdminController) AddMappedEntity(key string, entity repo.MappedEntity) {
	cntrl.mappedEntities[key] = entity
}

// Views all mapped entity.
const ADMIN_VIEW_ALL_ENTITIES_ROUTE = "admin view all entities route"

func (cntrl *AdminController) ViewAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityName := vars["entity"]

	entity, ok := cntrl.mappedEntities[entityName]
	if !ok {
		http.Error(w, "Entity type not found", http.StatusNotFound)
		return
	}

	query := fmt.Sprintf(`SELECT * FROM %s;`, entity.TableName())
	resultsMap, err := cntrl.repo.QueryToMap(query)
	if err != nil {
		cntrl.tmpl.SetMain(err.Error())
		cntrl.router.Response(w, cntrl.tmpl.Render())
		return
	}
	fmt.Println(resultsMap)


	// table := NewAdminTable()
	table := &AdminTable{DataMap: resultsMap}
	cntrl.tmpl.SetMain(table.Render())
	cntrl.router.Response(w, cntrl.tmpl.Render())

}

// Views mapped entity.
const ADMIN_VIEW_ENTITY_ROUTE = "admin view entity route"

func (cntrl *AdminController) ViewEntity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entityName := vars["entity"]
	uuid := vars["uuid"]

	mappedEntity, ok := cntrl.mappedEntities[entityName]
	if !ok {
		http.Error(w, "Entity type not found", http.StatusNotFound)
		return
	}

	// Retrieve mapped entity.
	err := cntrl.repo.GetByPrimaryKey(mappedEntity, uuid)
	if err != nil {
		cntrl.tmpl.SetMain(err.Error())
		cntrl.router.Response(w, cntrl.tmpl.Render())
		return
	}

	// Populate form.
	f, err := AdminForm(mappedEntity)
	if err != nil {
		cntrl.tmpl.SetMain(err.Error())
		cntrl.router.Response(w, cntrl.tmpl.Render())
		return
	}

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
	err := decoder.Decode(entityPointer, r.PostForm)
	if err != nil {
		cntrl.logger.Warn(err)
		cntrl.tmpl.SetMain(err.Error())
		cntrl.router.Response(w, cntrl.tmpl.Render())
		return
	}

	// Save entity.
	err = cntrl.repo.Save(entityPointer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		cntrl.logger.Warn(err)
		return
	}

	redirectUrl := cntrl.router.UrlInternal(ADMIN_VIEW_ENTITY_ROUTE, "entity", entityName, "uuid", uuid)
	http.Redirect(w, r, redirectUrl.Value, http.StatusFound)
}
