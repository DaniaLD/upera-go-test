package productDomain

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

type ProductRepository interface {
	Create(productData globalModel.Product) (product globalModel.Product, err error)
}