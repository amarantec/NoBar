package models

import "gorm.io/gorm"

type OrderItems struct {
	gorm.Model
	Orders     Orders
	OrdersID   uint `json:"orders_id" gorm:"not null"`
	Products   Products
	ProductsID uint    `json:"products_id" gorm:"not null"`
	Quantity   int64   `json:"quantity" gorm:"not null"`
	TotalPrice float64 `json:"total_price" gorm:"not null"`
}
