package productRevisionDomainservice

import (
	productRevisionDomain "github.com/DaniaLD/upera-go-test/internal/domain/product-revision"
)

type ProductRevisionService struct {
	productRevisionRepository productRevisionDomain.ProductRevisionRepository
}

func NewProductRevisionService(productRevisionRepository productRevisionDomain.ProductRevisionRepository) *ProductRevisionService {
	return &ProductRevisionService{
		productRevisionRepository: productRevisionRepository,
	}
}
