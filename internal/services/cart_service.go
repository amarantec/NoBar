package services

import (
	"context"
	"errors"
	"log"

	"github.com/amarantec/nobar/internal/models"
	"gorm.io/gorm"
)

func (s *ServicePostgres) GetCartProducts(ctx context.Context, customerId string) ([]models.CartProductResponse, error) {
	cartItems := []models.Carts{}
	cartProducts := []models.CartProductResponse{}

	if err :=
		s.Db.WithContext(ctx).
			Where("customer_id = ?", customerId).
			Find(&cartItems).Error; err != nil {
		return []models.CartProductResponse{}, err
	}

	for _, i := range cartItems {
		item := models.CartProductResponse{}
		product, _ := s.GetProduct(ctx, i.ProductsID)
		item.ProductsID = i.ProductsID
		item.Name = product.Name
		item.ImageURL = product.ImageURL
		item.Price = product.Price
		item.Quantity = i.Quantity

		cartProducts = append(cartProducts, item)
	}

	return cartProducts, nil
}

func (s *ServicePostgres) GetCartItemsCount(ctx context.Context, customerId string) (int64, error) {
	var count int64

	if err :=
		s.Db.WithContext(ctx).
			Model(&models.Carts{}).
			Where("customer_id = ?", customerId).
			Count(&count).Error; err != nil {
		count = 0
		return count, err
	}
	return count, nil
}

func (s *ServicePostgres) AddToCart(ctx context.Context, item models.Carts) (bool, error) {
	sameItem := models.Carts{}

	err := s.Db.WithContext(ctx).
		Where("customer_id = ? AND products_id = ?", item.CustomerID, item.ProductsID).
		First(&sameItem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		s.Db.WithContext(ctx).Create(&item)
		log.Println("Item not found. Successfully created.")
		return true, nil
	} else {
		sameItem.Quantity += item.Quantity
		s.Db.WithContext(ctx).Save(&sameItem)
		log.Println("Quantity successfully updated")
		return true, nil
	}
}

func (s *ServicePostgres) UpdateQuantity(ctx context.Context, customerId string, productsId uint, quantity int64) (bool, error) {
	if err :=
		s.Db.WithContext(ctx).
			Model(&models.Carts{}).
			Where("customer_id = ? AND products_id = ?", customerId, productsId).
			UpdateColumn("quantity", gorm.Expr("quantity + ?", quantity)).
			Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *ServicePostgres) RemoveItemFromCart(ctx context.Context, customerId string, productsId uint) (bool, error) {
	if err :=
		s.Db.WithContext(ctx).
			Where("customer_id = ? AND products_id = ?", customerId, productsId).
			Unscoped().Delete(&models.Carts{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
