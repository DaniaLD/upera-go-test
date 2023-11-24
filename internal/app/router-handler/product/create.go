package productRouterHandler

import (
	productDto "github.com/DaniaLD/upera-go-test/internal/app/router-handler/product/dto"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"github.com/gofiber/fiber/v2"
)

// CreateProduct
// @Summary CreateProduct
// @Description CreateProduct
// @Tags Product
// @Produce json
// @Param body body productDto.CreateProductRequestDto true "body"
// @Success 200 {object} globalModel.Product
// @Router /api/v1/product [post]
func (h *ProductRouterHandler) CreateProduct(c *fiber.Ctx) error {
	var body productDto.CreateProductRequestDto
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(globalModel.ErrorModel{Message: err.Error()})
	}

	product, err := h.productAppService.Create(globalModel.Product{
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
