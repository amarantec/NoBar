package models

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	CustomerID string `json:"customer_id" binding:"required" gorm:"not null"`
	ProductsID uint   `json:"products_id" binding:"required" gorm:"not null"`
	Quantity   int64  `json:"quantity" binding:"required" gorm:"not null"`
}
