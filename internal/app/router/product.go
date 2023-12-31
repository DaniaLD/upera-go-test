package router

import productRouterHandler "github.com/DaniaLD/upera-go-test/internal/app/router-handler/product"

func (r *Router) productRouter() {
	productGroup := r.v1.Group("product")

	productHandler := productRouterHandler.NewProductRouterHandler(*r.productAppService)
	productGroup.Post("/", productHandler.CreateProduct)
	productGroup.Put("/:id", productHandler.UpdateProduct)
	productGroup.Get("/:id", productHandler.GetProduct)
	productGroup.Get("/:id/revision/:revisionId", productHandler.GetProductByRevisionId)
}
