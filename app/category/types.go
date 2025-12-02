package category

import (
	"fmt"

	"github.com/mytheresa/go-hiring-challenge/models"
)

type Category struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Validate checks if the category fields are valid.
func (c *Category) Validate() error {
	if c.Code == "" {
		return fmt.Errorf("code is required")
	}
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

// ToModelCategory converts the API Category to a models.Category.
func (c *Category) ToModelCategory() *models.Category {
	return &models.Category{
		Code: c.Code,
		Name: c.Name,
	}
}

// FromModelCategory converts a models.Category to an API Category.
func FromModelCategory(m *models.Category) Category {
	return Category{
		Code: m.Code,
		Name: m.Name,
	}
}
