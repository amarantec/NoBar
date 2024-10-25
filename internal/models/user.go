package models

import (
    "gorm.io/gorm"
)

type Users struct {
    gorm.Model
    Name        string  `json:"name" binding:"required" gorm:"not null; unique"`
    Password    string  `json:"password" binding:"required" gorm:"not null"` 
}
