package productDomainService

import "github.com/DaniaLD/upera-go-test/internal/domain/product"

type ProductService struct {
	productRepository productDomain.ProductRepository
}

func NewProductService(productRepository productDomain.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}
