package productRevisionDomainservice

import (
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"github.com/DaniaLD/upera-go-test/utils"
	"strconv"
)

func (s *ProductRevisionService) Create(productId string, previousData globalModel.Product, newData globalModel.Product) {
	updatedAttributes := utils.MapsDiff(
		map[string]string{
			"name":        previousData.Name,
			"description": previousData.Description,
			"color":       previousData.Color,
			"price":       strconv.Itoa(int(previousData.Price)),
			"imageUrl":    previousData.ImageURL,
		},
		map[string]string{
			"name":        newData.Name,
			"description": newData.Description,
			"color":       newData.Color,
			"price":       strconv.Itoa(int(newData.Price)),
			"imageUrl":    newData.ImageURL,
		})
	if len(updatedAttributes) != 0 {
		s.productRevisionRepository.Create(productId, previousData, newData, updatedAttributes)
	}
	return
}
