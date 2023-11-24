package productAppService

import (
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
)

func (s *Service) Update(productData globalModel.Product) (globalModel.Product, error) {
	orgProduct, updateTime, err := s.productUseCase.Update(productData)
	if err != nil {
		return globalModel.Product{}, err
	}

	return globalModel.Product{
		Id:          productData.Id,
		Name:        productData.Name,
		Description: productData.Description,
		Color:       productData.Color,
		Price:       productData.Price,
		ImageURL:    productData.ImageURL,
		CreatedAt:   orgProduct.CreatedAt,
		UpdatedAt:   updateTime,
	}, nil
}
