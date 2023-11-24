package productRevisionDomainservice

import (
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
)

func (s *ProductRevisionService) Get(productId string) (revisions []globalModel.ProductRevision, err error) {
	return s.productRevisionRepository.Get(productId)
}
