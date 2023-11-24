package productRouterHandler

import (
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// GetProductByRevisionId
// @Summary GetProductByRevisionId
// @Description GetProductByRevisionId
// @Tags Product
// @Produce json
// @Param id path string true "Product Id" format(string)
// @Param revisionId path string true "Revision Id" format(string)
// @Success 200 {object} globalModel.Product
// @Router /api/v1/product/{id}/revision/{revisionId} [get]
func (h *ProductRouterHandler) GetProductByRevisionId(c *fiber.Ctx) error {
	prdId := c.Params("id")
	prdRvsId := c.Params("revisionId")
	prdRvsIdInt, err := strconv.ParseInt(prdRvsId, 10, 64)
	prd, err := h.productAppService.GetByRevisionId(prdId, prdRvsIdInt)
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
