package models

type Product struct {
	ProductId               int                      `gorm:"column:product_id;primary_key;auto_increment" json:"product_id"`
	ProductName             string                   `gorm:"column:product_name;size:255;not null" json:"product_name"`
	ProductDescription      string                   `gorm:"column:product_description;not null" json:"product_description"`
	ProductPrice            float64                  `gorm:"column:product_price;not null" json:"product_price"`
	CreatedAt               string                   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt               string                   `gorm:"column:updated_at;not null" json:"updated_at"`
	ProductImages           []ProductImage           `gorm:"foreignkey:ProductId" json:"product_images"`
	CompressedProductImages []CompressedProductImage `gorm:"foreignkey:ProductId" json:"compressed_product_images"`
}

func (p *Product) TableName() string {
	return "products"
}

type ProductImage struct {
	ProductId int    `gorm:"column:product_id;not null" json:"product_id"`
	ImageUrl  string `gorm:"column:image_url;not null" json:"image_url"`
}

func (p *ProductImage) TableName() string {
	return "product_images"
}

type CompressedProductImage struct {
	ProductId int    `gorm:"column:product_id;not null" json:"product_id"`
	ImageUrl  string `gorm:"column:image_url;not null" json:"image_url"`
}

func (p *CompressedProductImage) TableName() string {
	return "compressed_product_images"
}
