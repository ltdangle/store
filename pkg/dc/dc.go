package dc

import (
	"store/pkg/infra"
	"store/pkg/repo"
	"store/pkg/service"
	"store/pkg/web"

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

	CartController *web.CartController
}

func NewDc(envFile string ) *Dc {
	dc := &Dc{}
	cfg := infra.ReadConfig(envFile)
	dc.Db = infra.Gorm(cfg)

	dc.CustomerRepo = repo.NewCustomerRepo(dc.Db)
	dc.CustomerService = service.NewCustomerService(dc.CustomerRepo)

	dc.ProductRepo = repo.NewProductRepo(dc.Db)
	dc.ProductService = service.NewProductService(dc.ProductRepo, dc.Db)

	dc.CartRepo = repo.NewCartRepo(dc.Db)
	dc.CartService = service.NewCartService(dc.CartRepo)

	dc.CartController = web.NewCartController(dc.CartService,dc.CartRepo)
	return dc
}
