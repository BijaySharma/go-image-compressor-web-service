package controllers

import (
	"net/http"
	"strconv"
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

func (controller *ProductsController) AddCompressedImages(c *gin.Context) {
	logrus.Info("ProductsController.AddCompressedImages")

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error("Error converting product ID to integer: ", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDto{
			Status:  "error",
			Code:    400,
			Message: "Invalid product ID",
		})
	}

	var payload dto.AddCompressedImagesRequestDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		logrus.Error("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDto{
			Status:  "error",
			Code:    400,
			Message: "Error binding JSON",
		})

		return
	}

	res, apiErr := productsService.AddCompressedImages(productId, &payload)
	if apiErr != nil {
		logrus.Error("Error adding compressed images: ", err)
		c.JSON(apiErr.Code, apiErr)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (controller *ProductsController) GetAllProducts(c *gin.Context) {
	logrus.Info("ProductsController.GetProducts")

	res, err := productsService.GetAllProducts()
	if err != nil {
		logrus.Error("Error getting products: ", err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
