package productAppService

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *Service) GetByRevisionId(productId string, revisionId int64) (product globalModel.Product, err error) {
	return s.productRevisionUseCase.GetByRevisionId(productId, revisionId)
}
