package models

import "time"

type OrderOverviewResponse struct {
    OrderDate   time.Time
    TotalPrice  float64
    Product     string
    ImageURL    string
}
