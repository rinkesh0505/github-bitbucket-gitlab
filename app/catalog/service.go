package catalog

import (
	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/shopspring/decimal"
)

//go:generate mockgen -package=catalog -source=service.go -destination=service_mock.go
type CatalogService interface {
	GetProducts(req ListRequest) ([]Product, int64, error)
	GetProductDetails(code string) (*ProductDetails, error)
}

// catalogService handles business logic for the catalog.
type catalogService struct {
	repo models.ProductRepository
}

// NewCatalogService creates a new catalog service.
func NewCatalogService(repo models.ProductRepository) CatalogService {
	return &catalogService{
		repo: repo,
	}
}

// GetProducts fetches products for the given request and converts them to response format.
func (s *catalogService) GetProducts(req ListRequest) ([]Product, int64, error) {
	q := req.ToProductQuery()

	modelProducts, total, err := s.repo.GetProducts(q)
	if err != nil {
		return nil, 0, err
	}
	return mapToResponseProducts(modelProducts), total, nil
}

// GetProductDetails returns a product with its variants and category by product code.
// Variants without a specific price inherit the product price.
func (s *catalogService) GetProductDetails(code string) (*ProductDetails, error) {
	mp, err := s.repo.GetProductByCode(code)
	if err != nil {
		return nil, err
	}

	return mapProductToDetails(mp), nil
}

// mapToResponseProducts converts model products to API response products.
func mapToResponseProducts(modelProducts []models.Product) []Product {
	products := make([]Product, len(modelProducts))
	for i, p := range modelProducts {
		var cat *Category
		if p.Category != nil {
			cat = &Category{
				Code: p.Category.Code,
				Name: p.Category.Name,
			}
		}

		products[i] = Product{
			Code:     p.Code,
			Price:    priceToFloat64(p.Price),
			Category: cat,
		}
	}
	return products
}

// priceToFloat64 converts decimal price to float64.
func priceToFloat64(price decimal.Decimal) float64 {
	return price.InexactFloat64()
}

// mapProductToDetails converts a model Product to API ProductDetails.
func mapProductToDetails(mp *models.Product) *ProductDetails {
	pd := &ProductDetails{
		Code:  mp.Code,
		Price: priceToFloat64(mp.Price),
	}
	if mp.Category != nil {
		pd.Category = &Category{
			Code: mp.Category.Code,
			Name: mp.Category.Name,
		}
	}

	pd.Variants = make([]VariantDetails, len(mp.Variants))
	for i, v := range mp.Variants {
		pd.Variants[i] = mapVariantToDetails(v, mp.Price)
	}

	return pd
}

// mapVariantToDetails maps a model Variant to VariantDetails, inheriting
// the product price when the variant price is unspecified.
func mapVariantToDetails(v models.Variant, productPrice decimal.Decimal) VariantDetails {
	var price decimal.Decimal
	if v.Price.IsZero() {
		price = productPrice
	} else {
		price = v.Price
	}
	return VariantDetails{
		Name:  v.Name,
		SKU:   v.SKU,
		Price: priceToFloat64(price),
	}
}
