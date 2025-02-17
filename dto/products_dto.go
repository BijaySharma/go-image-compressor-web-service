package dto

import "web-service/models"

type CreateProductRequestDto struct {
	UserId             int      `json:"user_id" binding:"required"`
	ProductName        string   `json:"product_name" binding:"required"`
	ProductDescription string   `json:"product_description" binding:"required"`
	ProductPrice       float64  `json:"product_price" binding:"required"`
	ProductImages      []string `json:"product_images" binding:"required"`
}

type CreateProductResponseDto struct {
	Status  string         `json:"status"`
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Product models.Product `json:"product"`
}

type GetProductImagesResponseDto []string

type AddCompressedImagesRequestDto []string

type GetAllProductResponseDto struct {
	Status   string           `json:"status"`
	Code     int              `json:"code"`
	Message  string           `json:"message"`
	Products []models.Product `json:"products"`
}
