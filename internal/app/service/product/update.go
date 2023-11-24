package productAppService

import (
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
)

func (s *Service) Update(productData globalModel.Product) (globalModel.Product, error) {
	orgProduct, updateTime, err := s.productUseCase.Update(productData)
	if err != nil {
		return globalModel.Product{}, err
	}

	productData.CreatedAt = orgProduct.CreatedAt
	productData.UpdatedAt = updateTime
	go func() {
		s.productRevisionUseCase.Create(productData.Id, orgProduct, productData)
	}()

	return productData, nil
}
