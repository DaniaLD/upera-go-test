package productRevisionAppService

import (
	productRevisionDomain "github.com/DaniaLD/upera-go-test/internal/domain/product-revision"
)

type Service struct {
	productRevisionUseCase productRevisionDomain.ProductRevisionUseCase
}

func NewProductRevisionAppService(productRevisionUseCase productRevisionDomain.ProductRevisionUseCase) *Service {
	return &Service{
		productRevisionUseCase: productRevisionUseCase,
	}
}
