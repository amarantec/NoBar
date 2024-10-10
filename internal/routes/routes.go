package routes

import (
	"github.com/amarantec/nobar/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {

	r.POST("/welcome-customer", welcomeCustomer)

	categoriesGroup := r.Group("/categories")
	{
		categoriesGroup.POST("/insert-category", insertCategory)
		categoriesGroup.GET("/list-categories", listCategories)
		categoriesGroup.GET("/get-category/:categoryId", getCategory)
		categoriesGroup.DELETE("/delete-category/:categoryId", deleteCategory)
	}

	productsGroup := r.Group("/products")
	{
		productsGroup.POST("/insert-product", insertProduct)
		productsGroup.GET("/list-products", listProducts)
		productsGroup.GET("/get-product/:productId", getProduct)
		productsGroup.GET("/list-products-by-category/:categoryUrl", listProductsByCategory)
	}

	cartGroup := r.Group("/cart")
	{
		cartGroup.POST("/add-to-cart", middlewares.Auth, addToCart)
		cartGroup.PUT("/update-quantity/:productsId/:quantity", middlewares.Auth, updateQuantity)
		cartGroup.PUT("/remove-item-from-cart/:productsId", middlewares.Auth, removeItemFromCart)
		cartGroup.GET("/get-cart-items-count", middlewares.Auth, getCartItemsCount)
		cartGroup.GET("/get-cart-products", middlewares.Auth, getCartProducts)

	}

	orderGroup := r.Group("/order")
	{
		orderGroup.POST("/place-order", middlewares.Auth, placeOrder)
		orderGroup.GET("/get-order", middlewares.Auth, getOrders)
		orderGroup.GET("/get-order-details/:orderId", middlewares.Auth, getOrderDetails)
	}
}
