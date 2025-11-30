package catalog

import (
	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/shopspring/decimal"
)

// CatalogService handles business logic for the catalog.
type CatalogService struct {
	repo models.ProductRepository
}

// NewCatalogService creates a new catalog service.
func NewCatalogService(repo models.ProductRepository) *CatalogService {
	return &CatalogService{
		repo: repo,
	}
}

// GetAllProducts fetches all products and converts them to response format.
func (s *CatalogService) GetAllProducts() ([]Product, error) {
	modelProducts, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return mapToResponseProducts(modelProducts), nil
}

// mapToResponseProducts converts model products to API response products.
func mapToResponseProducts(modelProducts []models.Product) []Product {
	products := make([]Product, len(modelProducts))
	for i, p := range modelProducts {
		products[i] = Product{
			Code:  p.Code,
			Price: priceToFloat64(p.Price),
		}
	}
	return products
}

// priceToFloat64 converts decimal price to float64.
func priceToFloat64(price decimal.Decimal) float64 {
	return price.InexactFloat64()
}
