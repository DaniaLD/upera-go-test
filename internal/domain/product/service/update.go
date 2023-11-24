package productDomainService

import (
	"encoding/json"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
)

func (s *ProductService) Update(productData globalModel.Product) (globalModel.Product, error) {
	beforeUpdate, updateTime, err := s.productRepository.Update(productData)
	productData.CreatedAt = beforeUpdate.CreatedAt
	productData.UpdatedAt = updateTime
	go func() {
		message, _ := json.Marshal(globalModel.ProductUpdateEvent{
			ProductId:    productData.Id,
			PreviousData: beforeUpdate,
			NewData:      productData,
		})
		s.productRepository.DispatchUpdateEvent(message)
	}()
	return productData, err
}
