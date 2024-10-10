package models

import (
    "time"
    "gorm.io/gorm"
)

type Orders struct {
    gorm.Model
    CustomerID  string      `json:"customer_id" binding:"required" gorm:"not null"`
    OrderDate   time.Time
    TotalPrice  float64     `json:"total_price" gorm:"not null"`
    OrderItems  []OrderItems
}
