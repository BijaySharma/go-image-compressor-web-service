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
		// @METHOD: POST
		// @DESCRIPTION: Create a new product
		products.POST("/", productsController.CreateProduct)

		// @PATH: /products/images/:id
		// @METHOD: POST
		// @DESCRIPTION: Get All Product Images
		products.GET("images/:id", productsController.GetProductImages)
	}

}
