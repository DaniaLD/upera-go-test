package globalModel

type ProductRevision struct {
	Id                string   `json:"id"`
	ProductId         string   `json:"productId"`
	RevisionId        int64    `json:"revisionId"`
	UpdatedAttributes []string `json:"updatedAttributes"`
	PreviousValue     Product  `json:"previousValue"`
	NewValue          Product  `json:"newValue"`
	CreatedAt         string   `json:"createdAt"`
}
