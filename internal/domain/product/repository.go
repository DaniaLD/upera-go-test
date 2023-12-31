package productDomain

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

type ProductRepository interface {
	Create(productData globalModel.Product) (product globalModel.Product, err error)
	Update(productData globalModel.Product) (orgProduct globalModel.Product, updateTime string, err error)
	Get(productId string) (product globalModel.Product, err error)
	DispatchUpdateEvent(revisionData []byte)
}
