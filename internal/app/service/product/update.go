package productAppService

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *Service) Update(productData globalModel.Product) (product globalModel.Product, err error) {
	return s.productUseCase.Update(productData)
}
