package productRevisionAppService

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *Service) Create(productId string, previousData globalModel.Product, newData globalModel.Product) {
	s.productRevisionUseCase.Create(productId, previousData, newData)
}
