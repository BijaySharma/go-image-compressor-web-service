package services

import (
	"fmt"
	"strconv"
	"web-service/dto"
	"web-service/producer"
	"web-service/repository"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
		Product: *product,
	}, nil
}

func (service *ProductsService) GetProductImages(productId string) (*dto.GetProductImagesResponseDto, *dto.ErrorDto) {
	logrus.Info("ProductsService.GetProductImages")

	productImages, err := productsRepository.GetProductImages(productId)

	if err != nil {
		return nil, &dto.ErrorDto{
			Status:  "error",
			Code:    500,
			Message: err.Error(),
		}
	}

	return &productImages, nil
}

func (service *ProductsService) AddCompressedImages(productId int, payload *dto.AddCompressedImagesRequestDto) (*[]string, *dto.ErrorDto) {
	logrus.Info("ProductsService.AddCompressedImages")

	var imageUrls []string = make([]string, len(*payload))
	for i, imageUrl := range *payload {
		imageUrls[i] = fmt.Sprintf("%s/images/%s", viper.GetString("server.host"), imageUrl)
	}

	err := productsRepository.AddCompressedImages(productId, &imageUrls)

	if err != nil {
		return nil, &dto.ErrorDto{
			Status:  "error",
			Code:    500,
			Message: err.Error(),
		}
	}

	return &imageUrls, nil
}
