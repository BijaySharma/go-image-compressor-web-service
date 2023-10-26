package routes

import (
	"web-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProductsRoutes(router *gin.Engine) {
	productsController := controllers.ProductsController{}
	router.POST("/products", productsController.CreateProduct)
}
