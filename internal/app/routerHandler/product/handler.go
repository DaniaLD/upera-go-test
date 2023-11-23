package productRouterHandler

import productAppService "github.com/DaniaLD/upera-go-test/internal/app/service/product"

type ProductRouterHandler struct {
	productAppService productAppService.Service
}

func NewProductRouterHandler(productAppService productAppService.Service) *ProductRouterHandler {
	return &ProductRouterHandler{
		productAppService: productAppService,
	}
}
