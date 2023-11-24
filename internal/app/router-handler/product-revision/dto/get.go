package productRevisionDto

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

type ProductRevisionsListResponse struct {
	Revisions []globalModel.ProductRevision `json:"revisions"`
	Total     int64                         `json:"total"`
}
