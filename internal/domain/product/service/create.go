package productDomainService

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *ProductService) Create(productData globalModel.Product) (product globalModel.Product, err error) {
	return s.productRepository.Create(productData)
}
