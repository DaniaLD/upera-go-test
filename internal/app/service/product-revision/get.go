package productRevisionAppService

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *Service) Get(productId string) (revisions []globalModel.ProductRevision, err error) {
	return s.productRevisionUseCase.Get(productId)
}
