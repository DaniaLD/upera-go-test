package productDto

type CreateProductRequestDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Price       int64  `json:"price"`
	ImageURL    string `json:"imageURL"`
}
