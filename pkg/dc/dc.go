package dc

import (
	"store/pkg/infra"
	"store/pkg/logger"
	"store/pkg/repo"
	"store/pkg/service"
	"store/pkg/web"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Dependency container.
type Dc struct {
	Db              *gorm.DB
	CustomerRepo    *repo.CustomerRepo
	BaseProductRepo *repo.BaseProductRepo
	ProductRepo     *repo.ProductRepo
	CartRepo        *repo.CartRepo
	CartItemRepo    *repo.CartItemRepo

	CustomerService    *service.CustomerService
	BaseProductService *service.BaseProductService
	ProductService     *service.ProductService
	CartService        *service.CartService

	AppRouter *web.AppRouter
	Logger    logger.LoggerInterface

	CartController  *web.CartController
	AdminController *web.AdminController
}

func NewDc(envFile string) *Dc {
	dc := &Dc{}
	cfg := infra.ReadConfig(envFile)
	dc.Db = infra.Gorm(cfg.POSTGRES_URL)

	dc.Logger = logrus.New()

	dc.CustomerRepo = repo.NewCustomerRepo(dc.Db)
	dc.CustomerService = service.NewCustomerService(dc.CustomerRepo)

	dc.ProductRepo = repo.NewProductRepo(dc.Db)
	dc.ProductService = service.NewProductService(dc.ProductRepo, dc.Db)

	dc.CartRepo = repo.NewCartRepo(dc.Db)
	dc.CartItemRepo = repo.NewCartItemRepo(dc.Db)
	dc.CartService = service.NewCartService(dc.CartRepo, dc.CartItemRepo, dc.Db)

	dc.AppRouter = web.NewAppRouter(mux.NewRouter())

	tmpl := web.NewTmpl(dc.AppRouter)
	dc.CartController = web.NewCartController(dc.AppRouter, dc.CartService, dc.CartRepo, dc.Logger, tmpl, dc.Db)
	dc.AdminController = web.NewAdminController(dc.AppRouter, dc.Logger, tmpl, dc.Db)

	return dc
}
