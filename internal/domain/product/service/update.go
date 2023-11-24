package productDomainService

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *ProductService) Update(productData globalModel.Product) (orgProduct globalModel.Product, updateTime string, err error) {
	return s.productRepository.Update(productData)
}
