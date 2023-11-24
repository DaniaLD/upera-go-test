package productRevisionRouterHandler

import (
	productRevisionAppService "github.com/DaniaLD/upera-go-test/internal/app/service/product-revision"
)

type ProductRevisionRouterHandler struct {
	productRevisionAppService productRevisionAppService.Service
}

func NewProductRevisionRouterHandler(productRevisionAppService productRevisionAppService.Service) *ProductRevisionRouterHandler {
	return &ProductRevisionRouterHandler{
		productRevisionAppService: productRevisionAppService,
	}
}
