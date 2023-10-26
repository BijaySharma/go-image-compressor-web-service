package services

import (
	"web-service/dto"

	"github.com/sirupsen/logrus"
)

type ProductsService struct{}

func (service *ProductsService) CreateProduct(payload *dto.CreateProductRequestDto) (*dto.CreateProductResponseDto, *dto.ErrorDto) {
	logrus.Info("ProductsService.CreateProduct")
	return &dto.CreateProductResponseDto{
		Status:  "success",
		Code:    201,
		Message: "create product",
	}, nil
}
