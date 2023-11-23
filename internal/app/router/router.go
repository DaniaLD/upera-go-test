package router

import (
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	fiberApp    *fiber.App
	api         fiber.Router
	v1          fiber.Router
	fiberRouter fiber.Router
}

func NewRouter(
	fiberApp *fiber.App,
) *Router {
	return &Router{
		fiberApp: fiberApp,
	}
}

func (r *Router) InitRouter() {
	r.fiberRouter = r.fiberApp.Group("/")

	r.api = r.fiberApp.Group("/api")

	r.v1 = r.api.Group("/v1")

	r.swaggerRouter()
}
