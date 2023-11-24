package productRevisionAppService

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *Service) Get(productId string, limit int64, page int64) (revisions []globalModel.ProductRevision, total int64, err error) {
	skip := (page - 1) * limit
	return s.productRevisionUseCase.Get(productId, limit, skip)
}
