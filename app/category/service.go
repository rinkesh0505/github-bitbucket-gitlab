package category

import (
	"github.com/mytheresa/go-hiring-challenge/models"
)

type CategoryService struct {
	repo models.CategoryRepository
}

func NewCategoryService(repo models.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

// GetAllCategories returns all categories as API DTOs.
func (s *CategoryService) GetAllCategories() ([]Category, error) {
	cats, err := s.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return mapToCategories(cats), nil
}

// CreateCategory creates a new category in the DB.
func (s *CategoryService) CreateCategory(cat Category) error {
	return s.repo.CreateCategory(cat.ToModelCategory())
}

// mapToCategories converts model categories to API categories.
func mapToCategories(cats []models.Category) []Category {
	out := make([]Category, len(cats))
	for i, c := range cats {
		out[i] = FromModelCategory(&c)
	}
	return out
}
