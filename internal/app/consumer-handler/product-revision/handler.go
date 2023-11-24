package productRevisionConsumerHandler

import productRevisionAppService "github.com/DaniaLD/upera-go-test/internal/app/service/product-revision"

type ProductRevisionConsumerHandler struct {
	productRevisionAppService *productRevisionAppService.Service
}

func NewProductRevisionConsumerHandler(productRevisionAppService *productRevisionAppService.Service) *ProductRevisionConsumerHandler {
	return &ProductRevisionConsumerHandler{
		productRevisionAppService: productRevisionAppService,
	}
}
