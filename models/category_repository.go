package models

import (
	"gorm.io/gorm"
)

// CategoryRepository defines the interface for accessing category data.
type CategoryRepository interface {
	GetAllCategories() ([]Category, error)
	GetCategoryByID(id uint) (*Category, error)
	GetCategoryByCode(code string) (*Category, error)
	CreateCategory(category *Category) error
}

type CategoriesRepository struct {
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) *CategoriesRepository {
	return &CategoriesRepository{
		db: db,
	}
}

// GetAllCategories returns all categories.
func (r *CategoriesRepository) GetAllCategories() ([]Category, error) {
	var categories []Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoryByID returns a category by its ID.
func (r *CategoriesRepository) GetCategoryByID(id uint) (*Category, error) {
	var category Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// GetCategoryByCode returns a category by its code.
func (r *CategoriesRepository) GetCategoryByCode(code string) (*Category, error) {
	var category Category
	if err := r.db.Where("code = ?", code).First(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateCategory creates a new category.
func (r *CategoriesRepository) CreateCategory(category *Category) error {
	return r.db.Create(category).Error
}
