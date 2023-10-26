package routes

import (
	"web-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProductsRoutes(router *gin.Engine) {
	productsController := controllers.ProductsController{}
	products := router.Group("/products")
	{

		// @PATH: /products
		// @METHOD: GET
		// @DESCRIPTION: Get All Products
		products.GET("/", productsController.GetAllProducts)

		// @PATH: /products
		// @METHOD: POST
		// @DESCRIPTION: Create a new product
		products.POST("/", productsController.CreateProduct)

		// @PATH: /products/images/:id
		// @METHOD: GET
		// @DESCRIPTION: Get All Product Images
		products.GET("images/:id", productsController.GetProductImages)

		// @PATH: /products/compressed-images/:id
		// @METHOD: POST
		// @DESCRIPTION: Add Compressed Images
		products.POST("compressed-images/:id", productsController.AddCompressedImages)
	}

}
