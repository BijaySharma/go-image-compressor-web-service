package controllers

import "github.com/gin-gonic/gin"

type ProductsController struct{}

func (controller *ProductsController) CreateProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create product",
	})
}
