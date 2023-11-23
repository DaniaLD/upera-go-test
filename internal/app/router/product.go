package router

import productRouterHandler "github.com/DaniaLD/upera-go-test/internal/app/routerHandler/product"

func (r *Router) productRouter() {
	productGroup := r.v1.Group("product")

	productHandler := productRouterHandler.NewProductRouterHandler(*r.productAppService)
	productGroup.Post("/", productHandler.CreateProduct)
}
