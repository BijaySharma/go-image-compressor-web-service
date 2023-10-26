package controllers

import (
	"net/http"
	"web-service/dto"
	"web-service/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ProductsController struct{}

var productsService services.ProductsService = services.ProductsService{}

func (controller *ProductsController) CreateProduct(c *gin.Context) {
	logrus.Info("ProductsController.CreateProduct")

	var payload dto.CreateProductRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		logrus.Error("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDto{
			Status:  "error",
			Code:    400,
			Message: "Error binding JSON",
		})

		return
	}

	res, err := productsService.CreateProduct(&payload)
	if err != nil {
		logrus.Error("Error creating product: ", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (controller *ProductsController) GetProductImages(c *gin.Context) {
	logrus.Info("ProductsController.GetProductImages")
	productId := c.Param("id")

	res, err := productsService.GetProductImages(productId)
	if err != nil {
		logrus.Error("Error getting product images: ", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
