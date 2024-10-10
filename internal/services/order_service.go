package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/amarantec/nobar/internal/models"
	"gorm.io/gorm"
)

func (s *ServicePostgres) GetOrderDetails(ctx context.Context, customerId string, orderId uint) (models.OrderDetailsResponse, error) {
	order := models.Orders{}

	if err := s.Db.WithContext(ctx).
		Where("customer_id = ? AND id = ?", customerId, orderId).
		Preload("OrderItems.Products").
		Find(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.OrderDetailsResponse{}, nil
		}
		return models.OrderDetailsResponse{}, err
	}

	orderDetailsResponse := models.OrderDetailsResponse{
		OrderDate:  order.OrderDate,
		TotalPrice: order.TotalPrice,
		Products:   []models.OrderDetailsProductResponse{},
	}

	for _, item := range order.OrderItems {
		o := models.OrderDetailsProductResponse{}

		o.ProductsID = item.ProductsID
		o.Name = item.Products.Name
		o.ImageURL = item.Products.ImageURL
		o.Quantity = item.Quantity
		o.TotalPrice = item.TotalPrice

		orderDetailsResponse.Products = append(orderDetailsResponse.Products, o)
	}

	return orderDetailsResponse, nil

}

func (s *ServicePostgres) GetOrders(ctx context.Context, customerId string) ([]models.OrderOverviewResponse, error) {
	orders := []models.Orders{}

	if err :=
		s.Db.WithContext(ctx).
			Preload("OrderItems.Products").
			Where("customer_id = ?", customerId).
			Find(&orders).Error; err != nil {
		return []models.OrderOverviewResponse{}, err
	}

	orderResponse := []models.OrderOverviewResponse{}
	for _, item := range orders {
		o := models.OrderOverviewResponse{}
		o.OrderDate = item.OrderDate
		o.TotalPrice = item.TotalPrice

		if len(item.OrderItems) > 1 {
			o.Product = fmt.Sprintf("%s and %d more...", item.OrderItems[0].Products.Name, len(item.OrderItems)-1)
		} else if len(item.OrderItems) == 1 {
			o.Product = item.OrderItems[0].Products.Name
		} else {
			o.Product = "No Products"
		}

		if len(item.OrderItems) > 0 {
			o.ImageURL = item.OrderItems[0].Products.ImageURL
		}
		orderResponse = append(orderResponse, o)
	}

	return orderResponse, nil

}

func (s *ServicePostgres) PlaceOrder(ctx context.Context, customerId string) (bool, error) {
	var totalPrice float64

	products, _ := s.GetCartProducts(ctx, customerId)

	for _, v := range products {
		totalPrice += v.Price * float64(v.Quantity)
	}

	orderItems := []models.OrderItems{}
	for _, items := range products {
		o := models.OrderItems{}
		o.ProductsID = items.ProductsID
		o.Quantity = items.Quantity
		o.TotalPrice = items.Price * float64(items.Quantity)

		orderItems = append(orderItems, o)
	}

	order := models.Orders{
		CustomerID: customerId,
		OrderDate:  time.Now(),
		TotalPrice: totalPrice,
		OrderItems: orderItems,
	}

	if err :=
		s.Db.WithContext(ctx).
			Create(&order).Error; err != nil {
		return false, err
	}

	if err :=
		s.Db.WithContext(ctx).
			Where("customer_id = ?", customerId).
			Unscoped().Delete(&models.Carts{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
