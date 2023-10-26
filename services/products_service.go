package services

import (
	"strconv"
	"web-service/dto"
	"web-service/producer"
	"web-service/repository"

	"github.com/sirupsen/logrus"
)

type ProductsService struct{}

var productsRepository repository.ProductsRepository = repository.ProductsRepository{}

func (service *ProductsService) CreateProduct(payload *dto.CreateProductRequestDto) (*dto.CreateProductResponseDto, *dto.ErrorDto) {
	logrus.Info("ProductsService.CreateProduct")

	product, err := productsRepository.CreateProduct(payload)

	if err != nil {
		return nil, &dto.ErrorDto{
			Status:  "error",
			Code:    500,
			Message: err.Error(),
		}
	}
	logrus.Info("Product created successfully with ID ", product.ProductId)

	// Publish message to Kafka
	producer.PublishMessage("products", strconv.Itoa(product.ProductId))

	return &dto.CreateProductResponseDto{
		Status:  "success",
		Code:    201,
		Message: "Product created successfully",
	}, nil
}
