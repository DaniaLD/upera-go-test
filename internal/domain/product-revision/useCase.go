package productRevisionDomain

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

type ProductRevisionUseCase interface {
	Create(productId string, previousData globalModel.Product, newData globalModel.Product)
}
