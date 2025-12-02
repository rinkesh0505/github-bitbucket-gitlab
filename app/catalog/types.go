package catalog

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/shopspring/decimal"
)

// Response represents the API response for the catalog endpoint.
type Response struct {
	Products []Product `json:"products"`
	Total    int64     `json:"total"`
}

// Product represents a product in the catalog API response.
type Product struct {
	Code     string    `json:"code"`
	Price    float64   `json:"price"`
	Category *Category `json:"category,omitempty"`
}

// Category represents a product category in the API response.
type Category struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// ListRequest represents pagination and filter parameters for listing products.
type ListRequest struct {
	Offset   int      `json:"offset"`
	Limit    int      `json:"limit"`
	Category string   `json:"category,omitempty"`
	PriceLT  *float64 `json:"price_lt,omitempty"`
}

// Validate adjusts and clamps request values to acceptable ranges.
// It applies defaults: Offset=0, Limit=10; and clamps Limit to [1,100].
func (r *ListRequest) Validate() {
	if r.Offset < 0 {
		r.Offset = 0
	}
	if r.Limit == 0 {
		r.Limit = 10
	}
	if r.Limit < 1 {
		r.Limit = 1
	}
	if r.Limit > 100 {
		r.Limit = 100
	}
}

// NewListRequestFromValues parses pagination params from url.Values and
// returns a validated ListRequest. Returns an error if parameters are malformed.
func NewListRequestFromValues(q url.Values) (ListRequest, error) {
	req := ListRequest{}

	if v := q.Get("offset"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil {
			return req, fmt.Errorf("invalid offset")
		}
		req.Offset = n
	}

	if v := q.Get("limit"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil {
			return req, fmt.Errorf("invalid limit")
		}
		req.Limit = n
	}

	if v := q.Get("category"); v != "" {
		req.Category = v
	}

	if v := q.Get("price_lt"); v != "" {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return req, fmt.Errorf("invalid price_lt")
		}
		req.PriceLT = &f
	}

	req.Validate()
	return req, nil
}

// ToProductQuery converts the catalog ListRequest into a models.ProductQuery
func (r ListRequest) ToProductQuery() models.ProductQuery {
	var priceDecimal *decimal.Decimal
	if r.PriceLT != nil {
		d := decimal.NewFromFloat(*r.PriceLT)
		priceDecimal = &d
	}
	return models.ProductQuery{
		Offset:   r.Offset,
		Limit:    r.Limit,
		Category: r.Category,
		PriceLT:  priceDecimal,
	}
}

// ProductDetails represents product response including variants.
type ProductDetails struct {
	Code     string           `json:"code"`
	Price    float64          `json:"price"`
	Category *Category        `json:"category,omitempty"`
	Variants []VariantDetails `json:"variants"`
}

type VariantDetails struct {
	Name  string  `json:"name"`
	SKU   string  `json:"sku"`
	Price float64 `json:"price"`
}
