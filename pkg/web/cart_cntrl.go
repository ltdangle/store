package web

import (
	"fmt"
	"net/http"
	"store/pkg/logger"
	"store/pkg/models"
	"store/pkg/repo"
	"store/pkg/service"
	"store/pkg/web/form"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CartController struct {
	router  *mux.Router
	service *service.CartService
	repo    *repo.CartRepo
	logger  logger.LoggerInterface
	tmpl    *Tmpl
	db      *gorm.DB
}

func NewCartController(router *mux.Router, service *service.CartService, repo *repo.CartRepo, logger logger.LoggerInterface, tmpl *Tmpl, db *gorm.DB) *CartController {
	return &CartController{router: router, service: service, repo: repo, logger: logger, tmpl: tmpl, db: db}
}

type CartVM struct {
	Cart *models.Cart
}

const CART_VIEW_ROUTE = "cart"

func (cntrl *CartController) View(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	cart, err := cntrl.repo.FindByUuid(uuid)

	if err != nil {
		cntrl.logger.Warn(fmt.Sprintf("CartController.View: cart %s : %s", uuid, err.Error()))
		fmt.Fprint(w, err.Error())
	} else {
		vm := CartVM{
			Cart: cart,
		}
		cartTmpl := NewCartTmpl(cntrl.router)
		cntrl.tmpl.setMain(cartTmpl.cart(vm))
		response(w, cntrl.tmpl.render())
	}
}

const CART_ITEM_DELETE_ROUTE = "delete cart item"

func (cntrl *CartController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemUuid := vars["uuid"]

	cart, err := cntrl.service.RemoveCartItem(cartItemUuid)
	if err != nil {
		cntrl.logger.Warn(fmt.Sprintf("CartController.DeleteItem: cart with cartItem %s : %s", cartItemUuid, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cartUrl := UrlInternal(cntrl.router, CART_VIEW_ROUTE, "uuid", cart.Uuid)
	if cartUrl.Error != nil {
		cntrl.logger.Warn("CartController.DeleteItem: could not parse redirect url")
	}

	cntrl.logger.Info(fmt.Sprintf("CartController.DeleteItem: item deleted, redirect to %s", cartUrl))

	http.Redirect(w, r, cartUrl.Value, http.StatusTemporaryRedirect)
}

const CART_ITEM_EDIT_ROUTE = "edit cart item"

func (cntrl *CartController) EditCartItem(w http.ResponseWriter, r *http.Request) {
	var columnNames []string

	// List columns of a model's table
	columns, err := cntrl.db.Migrator().ColumnTypes(&models.Cart{})

	if err == nil {
		for _, column := range columns {
			columnNames = append(columnNames, column.Name()+" - "+column.DatabaseTypeName()+"-"+column.ScanType().String()+";")
		}
	}
	fmt.Println(columnNames)

	// Form.
	// f := form.NewForm()
	// f.Method = "POST"
	// f.Action = UrlInternal(cntrl.router, CART_ITEM_DELETE_ROUTE).Value
	// f.AddField(&form.Field{Name: "Text", Type: "text", Value: "text value", Required: true})
	// f.AddField(&form.Field{Name: "Number", Type: "number", Value: "", Required: true})
	// f.AddField(&form.Field{Name: "Email", Type: "email", Value: "", Required: true, Error: "dis is incorrect"})
	// f.AddField(&form.Field{Name: "Password", Type: "password", Value: "", Required: true})
	// f.AddField(&form.Field{Name: "Date", Type: "date", Value: "", Required: true})
	// f.AddField(&form.Field{Name: "File", Type: "file", Value: "", Required: true})

	f := form.GormToForm(&models.Cart{}, cntrl.db)
	cntrl.tmpl.setMain(f.Render())
	response(w, cntrl.tmpl.render())
}
