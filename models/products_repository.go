package models

import (
	"gorm.io/gorm"
)

// ProductRepository defines the interface for fetching products with pagination and filters.
type ProductRepository interface {
	GetProducts(query ProductQuery) ([]Product, int64, error)
	GetProductByCode(code string) (*Product, error)
}

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{
		db: db,
	}
}

func (r *ProductsRepository) GetProducts(query ProductQuery) ([]Product, int64, error) {
	var products []Product
	var total int64

	base := r.db.Model(&Product{})
	if query.Category != "" {
		base = base.Joins("JOIN categories ON categories.id = products.category_id").Where("categories.code = ?", query.Category)
	}
	if query.PriceLT != nil {
		base = base.Where("price < ?", query.PriceLT)
	}

	if err := base.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := base.Preload("Variants").Preload("Category").Offset(query.Offset).Limit(query.Limit).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// GetProductByCode finds a product by its code and preloads variants and category.
func (r *ProductsRepository) GetProductByCode(code string) (*Product, error) {
	var p Product
	if err := r.db.Preload("Variants").Preload("Category").Where("code = ?", code).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
