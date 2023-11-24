package productAppService

import (
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
)

func (s *Service) Update(productData globalModel.Product) (globalModel.Product, error) {
	updatedProduct, err := s.productUseCase.Update(productData)
	if err != nil {
		return globalModel.Product{}, err
	}

	return updatedProduct, nil
}
