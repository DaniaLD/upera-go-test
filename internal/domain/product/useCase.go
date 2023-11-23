package productDomain

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

type ProductUseCase interface {
	Create(productData globalModel.Product) (product globalModel.Product, err error)
	Update(productData globalModel.Product) (product globalModel.Product, err error)
}
