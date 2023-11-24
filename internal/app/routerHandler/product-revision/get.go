package productRevisionRouterHandler

import (
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"github.com/gofiber/fiber/v2"
)

// GetProductRevisions
// @Summary GetProductRevisions
// @Description GetProductRevisions
// @Tags Product Revision
// @Produce json
// @Param productId path string true "Product Id" format(string)
// @Success 200 {object} []globalModel.ProductRevision
// @Router /api/v1/product-revision/{productId} [get]
func (h *ProductRevisionRouterHandler) GetProductRevisions(c *fiber.Ctx) error {
	prdId := c.Params("productId")
	revisions, err := h.productRevisionAppService.Get(prdId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(globalModel.ErrorModel{Message: err.Error()})
	} else if len(revisions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(globalModel.ResponseModel{
			Status:  404,
			Message: "not found",
			Payload: []globalModel.ProductRevision{},
		})
	}
	return c.JSON(revisions)
}
