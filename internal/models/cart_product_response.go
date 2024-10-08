package models

type CartProductResponse struct {
	ProductsID uint    `json:"product_id"`
	Name      string  `json:"name"`
	ImageURL  string  `json:"image_url"`
	Price     float64 `json:"price"`
	Quantity  int64   `json:"quantity"`
}
