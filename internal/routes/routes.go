package routes

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine) {
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
	}
}
