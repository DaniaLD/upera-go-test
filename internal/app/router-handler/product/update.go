package productRouterHandler

import (
	productDto "github.com/DaniaLD/upera-go-test/internal/app/router-handler/product/dto"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"github.com/gofiber/fiber/v2"
)

// UpdateProduct
// @Summary UpdateProduct
// @Description UpdateProduct
// @Tags Product
// @Produce json
// @Param body body productDto.CreateProductRequestDto true "body"
// @Param id path string true "Product Id" format(string)
// @Success 200 {object} productDto.ProductCommonResponseDto
// @Router /api/v1/product/{id} [put]
func (h *ProductRouterHandler) UpdateProduct(c *fiber.Ctx) error {
	var body productDto.CreateProductRequestDto
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(globalModel.ErrorModel{Message: err.Error()})
	}
	prdId := c.Params("id")

	product, err := h.productAppService.Update(globalModel.Product{
		Id:          prdId,
		Name:        body.Name,
		Description: body.Description,
		Color:       body.Color,
		Price:       body.Price,
		ImageURL:    body.ImageURL,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(globalModel.ErrorModel{Message: err.Error()})
	}
	return c.JSON(product)
}
