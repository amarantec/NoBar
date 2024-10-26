package services

import (
	"context"

	"github.com/amarantec/nobar/internal/models"
	"gorm.io/gorm"
)

type Service interface {
	WelcomeCustomer(ctx context.Context, customer models.Customer) (bool, error)

    Register (ctx context.Context, user models.Users) (bool, error)
    Login (ctx context.Context, user models.Users) (uint, error)

	InsertCategory(ctx context.Context, category models.Categories) (models.Categories, error)
	ListCategories(ctx context.Context) ([]models.Categories, error)
	GetCategory(ctx context.Context, id uint) (models.Categories, error)
    UpdateCategory(ctx context.Context, category models.Categories) (bool, error)
	DeleteCategory(ctx context.Context, id uint) (bool, error)

	InsertProduct(ctx context.Context, product models.Products) (models.Products, error)
	ListProducts(ctx context.Context) ([]models.Products, error)
	GetProduct(ctx context.Context, id uint) (models.Products, error)
	ListProductsByCategory(ctx context.Context, categoryUrl string) ([]models.Products, error)
    SearchProducts(ctx context.Context, searchText string) ([]models.Products, error)
    UpdateProduct(ctx context.Context, product models.Products) (bool, error)
    DeleteProduct(ctx context.Context, id uint) (bool, error)

	GetCartProducts(ctx context.Context, customerId string) ([]models.CartProductResponse, error)
	GetCartItemsCount(ctx context.Context, customerId string) (int64, error)
	AddToCart(ctx context.Context, item models.Carts) (bool, error)
	UpdateQuantity(ctx context.Context, customerId string, productsId uint, quantity int64) (bool, error)
	RemoveItemFromCart(ctx context.Context, customerId string, productsId uint) (bool, error)

	GetOrderDetails(ctx context.Context, customerId string, orderId uint) (models.OrderDetailsResponse, error)
	GetOrders(ctx context.Context, customerId string) ([]models.OrderOverviewResponse, error)
	PlaceOrder(ctx context.Context, customerId string) (bool, error)
}

type ServicePostgres struct {
	Db *gorm.DB
}
