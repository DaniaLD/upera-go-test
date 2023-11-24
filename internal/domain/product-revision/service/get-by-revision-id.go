package productRevisionDomainservice

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

func (s *ProductRevisionService) GetByRevisionId(productId string, revisionId int64) (product globalModel.Product, err error) {
	return s.productRevisionRepository.GetByRevisionId(productId, revisionId)
}
