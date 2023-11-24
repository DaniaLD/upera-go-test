package productAppService

import (
	productDomain "github.com/DaniaLD/upera-go-test/internal/domain/product"
	productRevisionDomain "github.com/DaniaLD/upera-go-test/internal/domain/product-revision"
)

type Service struct {
	productUseCase         productDomain.ProductUseCase
	productRevisionUseCase productRevisionDomain.ProductRevisionUseCase
}

func NewProductAppService(productUseCase productDomain.ProductUseCase, productRevisionUseCase productRevisionDomain.ProductRevisionUseCase) *Service {
	return &Service{
		productUseCase:         productUseCase,
		productRevisionUseCase: productRevisionUseCase,
	}
}
