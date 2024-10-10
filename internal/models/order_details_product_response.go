package models

type OrderDetailsProductResponse struct {
	ProductsID uint
	Name       string
	ImageURL   string
	Quantity   int64
	TotalPrice float64
}
