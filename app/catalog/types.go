package catalog

// Response represents the API response for the catalog endpoint.
type Response struct {
	Products []Product `json:"products"`
}

// Product represents a product in the catalog API response.
type Product struct {
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}
