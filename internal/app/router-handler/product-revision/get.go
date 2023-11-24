package productRevisionRouterHandler

import (
	productRevisionDto "github.com/DaniaLD/upera-go-test/internal/app/router-handler/product-revision/dto"
	globalModel "github.com/DaniaLD/upera-go-test/pkg/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// GetProductRevisions
// @Summary GetProductRevisions
// @Description GetProductRevisions
// @Tags Product Revision
// @Produce json
// @Param productId path string true "Product Id" format(string)
// @Param limit query int false "limit" default(10)
// @Param page query int false "page" default(1)
// @Success 200 {object} productRevisionDto.ProductRevisionsListResponse
// @Router /api/v1/product-revision/{productId} [get]
func (h *ProductRevisionRouterHandler) GetProductRevisions(c *fiber.Ctx) error {
	prdId := c.Params("productId")
	limit := c.Query("limit", "10")
	limitInt, _ := strconv.ParseInt(limit, 10, 64)
	page := c.Query("page", "1")
	pageInt, _ := strconv.ParseInt(page, 10, 64)
	revisions, total, err := h.productRevisionAppService.Get(prdId, limitInt, pageInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(globalModel.ErrorModel{Message: err.Error()})
	} else if len(revisions) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(globalModel.ResponseModel{
			Status:  404,
			Message: "not found",
			Payload: []globalModel.ProductRevision{},
		})
	}
	return c.JSON(productRevisionDto.ProductRevisionsListResponse{Revisions: revisions, Total: total})
}
