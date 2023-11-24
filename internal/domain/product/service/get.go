package productDomainService

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *ProductService) Get(productId string) (product globalModel.Product, err error) {
	return s.productRepository.Get(productId)
}
