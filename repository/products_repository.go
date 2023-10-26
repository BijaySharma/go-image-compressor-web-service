package repository

import (
	"fmt"
	"time"
	"web-service/db"
	"web-service/dto"
	"web-service/models"

	"github.com/sirupsen/logrus"
)

type ProductsRepository struct{}

func (repository *ProductsRepository) CreateProduct(payload *dto.CreateProductRequestDto) (*models.Product, error) {
	logrus.Info("ProductsRepository.CreateProduct")

	conn := db.GetInstance()
	tx := conn.Begin()
	defer tx.Rollback()

	product := models.Product{
		ProductName:        payload.ProductName,
		ProductPrice:       payload.ProductPrice,
		ProductDescription: payload.ProductDescription,
		CreatedAt:          time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:          time.Now().Format("2006-01-02 15:04:05"),
	}

	if err := tx.Create(&product).Error; err != nil {
		logrus.Error("Error creating product: ", err)
		return nil, err
	}

	var productImages []models.ProductImage = make([]models.ProductImage, len(payload.ProductImages))
	for i, image := range payload.ProductImages {
		productImages[i] = models.ProductImage{
			ProductId: product.ProductId,
			ImageUrl:  image,
		}
	}

	if err := tx.Create(&productImages).Error; err != nil {
		logrus.Error("Error creating product images: ", err)
		return nil, err
	}
	product.ProductImages = productImages

	if err := tx.Commit().Error; err != nil {
		logrus.Error("Error committing transaction: ", err)
		return nil, fmt.Errorf("Error committing transaction: %v", err)
	}

	return &product, nil
}

func (repository *ProductsRepository) GetProductImages(productId string) (dto.GetProductImagesResponseDto, error) {
	logrus.Info("ProductsRepository.GetProductImages")

	conn := db.GetInstance()
	var productImages []string

	if err := conn.Model(&models.ProductImage{}).Select("image_url").Where("product_id = ?", productId).Find(&productImages).Error; err != nil {
		logrus.Error("Error getting product images: ", err)
		return nil, err
	}

	return productImages, nil
}

func (repository *ProductsRepository) AddCompressedImages(productId int, imageUrls *[]string) error {
	logrus.Info("ProductsRepository.AddCompressedImages")

	conn := db.GetInstance()

	var compressedProductImages []models.CompressedProductImage = make([]models.CompressedProductImage, len(*imageUrls))

	for i, imageUrl := range *imageUrls {
		compressedProductImages[i] = models.CompressedProductImage{
			ProductId: productId,
			ImageUrl:  imageUrl,
		}
	}

	if err := conn.Create(&compressedProductImages).Error; err != nil {
		logrus.Error("Error creating product images: ", err)
		return err
	}

	return nil
}
