package globalModel

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Price       int64  `json:"price"`
	ImageURL    string `json:"imageURL"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type ProductUpdateEvent struct {
	ProductId    string  `json:"productId"`
	PreviousData Product `json:"previousData"`
	NewData      Product `json:"newData"`
}
