package productDto

import globalModel "github.com/DaniaLD/upera-go-test/pkg/model"

type CreateProductRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Price       int64  `json:"price"`
	ImageURL    string `json:"imageURL"`
}

type ProductCommonResponseDto struct {
	globalModel.ResponseModel
	Payload globalModel.Product `json:"payload"`
}
