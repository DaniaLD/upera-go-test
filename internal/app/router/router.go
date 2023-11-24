package router

import (
	productAppService "github.com/DaniaLD/upera-go-test/internal/app/service/product"
	productRevisionAppService "github.com/DaniaLD/upera-go-test/internal/app/service/product-revision"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	fiberApp                  *fiber.App
	api                       fiber.Router
	v1                        fiber.Router
	fiberRouter               fiber.Router
	productAppService         *productAppService.Service
	productRevisionAppService *productRevisionAppService.Service
}

func NewRouter(
	fiberApp *fiber.App,
	productAppService *productAppService.Service,
	productRevisionAppService *productRevisionAppService.Service,
) *Router {
	return &Router{
		fiberApp:                  fiberApp,
		productAppService:         productAppService,
		productRevisionAppService: productRevisionAppService,
	}
}

func (r *Router) InitRouter() {
	r.fiberRouter = r.fiberApp.Group("/")

	r.api = r.fiberApp.Group("/api")

	r.v1 = r.api.Group("/v1")

	r.swaggerRouter()
	r.productRouter()
	r.productRevisionRouter()
}
