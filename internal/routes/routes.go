package routes

import (
	"github.com/amarantec/nobar/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) {
	r.GET("/home", homePage)

	r.POST("/welcome-customer", welcomeCustomer)


    userGroup := r.Group("/user")
    {
        userGroup.POST("/register", register)
        userGroup.POST("/login", login)
    }

	categoriesGroup := r.Group("/categories")
	{
		categoriesGroup.POST("/insert-category", middlewares.Auth, insertCategory)
		categoriesGroup.GET("/list-categories", listCategories)
		categoriesGroup.GET("/get-category/:categoryId", getCategory)
        categoriesGroup.PUT("/update-category", middlewares.Auth, updateCategory)
		categoriesGroup.DELETE("/delete-category/:categoryId", middlewares.Auth, deleteCategory)
	}

	productsGroup := r.Group("/products")
	{
		productsGroup.POST("/insert-product", middlewares.Auth, insertProduct)
		productsGroup.GET("/list-products", listProducts)
		productsGroup.GET("/get-product/:productId", getProduct)
		productsGroup.GET("/list-products-by-category/:categoryUrl", listProductsByCategory)
        productsGroup.GET("/search-products/:searchText", searchProducts)
        productsGroup.PUT("/update-product", middlewares.Auth, updateProduct)
        productsGroup.DELETE("/delete-product/:productId", middlewares.Auth, deleteProduct)
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
