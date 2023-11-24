package productRouterHandler

import (
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"github.com/gofiber/fiber/v2"
)

// GetProduct
// @Summary GetProduct
// @Description GetProduct
// @Tags Product
// @Produce json
// @Param id path string true "Product Id" format(string)
// @Success 200 {object} globalModel.Product
// @Router /api/v1/product/{id} [get]
func (h *ProductRouterHandler) GetProduct(c *fiber.Ctx) error {
	prdId := c.Params("id")
	prd, err := h.productAppService.Get(prdId)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return c.Status(fiber.StatusNotFound).JSON(globalModel.ResponseModel{
			Status:  404,
			Message: "not found",
			Payload: map[string]string{},
		})
	} else if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(globalModel.ErrorModel{Message: err.Error()})
	}
	return c.JSON(prd)
}
