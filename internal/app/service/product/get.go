package productAppService

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *Service) Get(productId string) (product globalModel.Product, err error) {
	return s.productUseCase.Get(productId)
}
