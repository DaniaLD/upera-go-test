package productDomainService

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *ProductService) Update(productData globalModel.Product) (product globalModel.Product, err error) {
	return s.productRepository.Update(productData)
}
