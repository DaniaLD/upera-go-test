package productRevisionDomain

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

type ProductRevisionUseCase interface {
	Create(productId string, previousData globalModel.Product, newData globalModel.Product)
	Get(productId string, limit int64, skip int64) (revisions []globalModel.ProductRevision, total int64, err error)
	GetByRevisionId(productId string, revisionId int64) (product globalModel.Product, err error)
}
