package productAppService

import productDomain "github.com/DaniaLD/upera-go-test/internal/domain/product"

type Service struct {
	productUseCase productDomain.ProductUseCase
}

func NewProductAppService(productUseCase productDomain.ProductUseCase) *Service {
	return &Service{
		productUseCase: productUseCase,
	}
}
