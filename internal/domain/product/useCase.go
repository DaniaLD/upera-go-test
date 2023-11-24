package productDomain

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

type ProductUseCase interface {
	Create(productData globalModel.Product) (product globalModel.Product, err error)
	Update(productData globalModel.Product) (updatedProduct globalModel.Product, err error)
	Get(productId string) (product globalModel.Product, err error)
}
