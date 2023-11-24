package productRevisionDomain

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

type ProductRevisionRepository interface {
	Create(productId string, previousData globalModel.Product, newData globalModel.Product, updatedAttributes []string) error
	Get(productId string) (revisions []globalModel.ProductRevision, err error)
}
