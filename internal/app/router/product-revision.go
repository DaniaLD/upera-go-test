package router

import productRevisionRouterHandler "github.com/DaniaLD/upera-go-test/internal/app/router-handler/product-revision"

func (r *Router) productRevisionRouter() {
	prdRvsGroup := r.v1.Group("product-revision")

	prdRvsHandler := productRevisionRouterHandler.NewProductRevisionRouterHandler(*r.productRevisionAppService)
	prdRvsGroup.Get("/:productId", prdRvsHandler.GetProductRevisions)
}
