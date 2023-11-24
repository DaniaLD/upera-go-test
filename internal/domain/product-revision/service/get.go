package productRevisionDomainservice

import (
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
)

func (s *ProductRevisionService) Get(productId string, limit int64, skip int64) (revisions []globalModel.ProductRevision, total int64, err error) {
	return s.productRevisionRepository.Get(productId, limit, skip)
}
