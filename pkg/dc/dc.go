package dc

import (
	"store/pkg/infra"
	"store/pkg/logger"
	"store/pkg/repo"
	"store/pkg/service"
	"store/pkg/web"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Dependency container.
type Dc struct {
	Db           *gorm.DB
	Sqlx         *sqlx.DB
	GeneralRepo  *repo.GeneralRepo
	CustomerRepo *repo.CustomerRepo
	ProductRepo  *repo.ProductRepo
	CartRepo     *repo.CartRepo
	CartItemRepo *repo.CartItemRepo

	CustomerService *service.CustomerService
	ProductService  *service.ProductService
	CartService     *service.CartService

	AppRouter *web.AppRouter
	Logger    logger.LoggerInterface

	CartController  *web.CartController
	AdminController *web.AdminController

	AdminTemplate *web.AdminTmpl
}

func NewDc(envFile string) *Dc {
	dc := &Dc{}
	cfg := infra.ReadConfig(envFile)

	dc.Db = infra.Gorm(cfg.POSTGRES_URL)
	dc.Sqlx = infra.Sqlx(cfg.POSTGRES_URL)

	dc.Logger = logrus.New()

	dc.GeneralRepo = repo.NewGeneralRepo(dc.Sqlx, dc.Db)
	dc.CustomerRepo = repo.NewCustomerRepo(dc.Db)
	dc.CustomerService = service.NewCustomerService(dc.CustomerRepo)

	dc.ProductRepo = repo.NewProductRepo(dc.Db)
	dc.ProductService = service.NewProductService(dc.ProductRepo, dc.Db)

	dc.CartRepo = repo.NewCartRepo(dc.Db, dc.Sqlx)
	dc.CartItemRepo = repo.NewCartItemRepo(dc.Db)
	dc.CartService = service.NewCartService(dc.CartRepo, dc.CartItemRepo, dc.Db)

	dc.AppRouter = web.NewAppRouter(mux.NewRouter(), dc.Logger)

	dc.AdminTemplate = web.NewAdminTempl(dc.AppRouter)

	dc.CartController = web.NewCartController(dc.AppRouter, dc.CartService, dc.CartRepo, dc.Logger, dc.AdminTemplate, dc.Db)
	dc.AdminController = web.NewAdminController(dc.AppRouter, dc.Logger, dc.AdminTemplate, dc.GeneralRepo)

	return dc
}
