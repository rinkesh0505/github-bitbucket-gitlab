package catalog

import (
	"encoding/json"
	"net/http"
	"strings"
)

type CatalogHandler struct {
	service CatalogService
}

func NewCatalogHandler(service CatalogService) *CatalogHandler {
	return &CatalogHandler{
		service: service,
	}
}

// HandleGet supports offset/limit pagination via query params offset and limit.
// Defaults: offset=0, limit=10. Limit is clamped to [1,100].
func (h *CatalogHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	req, err := NewListRequestFromValues(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	products, total, err := h.service.GetProducts(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Products: products,
		Total:    total,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleGetDetails returns full product details.
func (h *CatalogHandler) HandleGetDetails(w http.ResponseWriter, r *http.Request) {
	// Expect the code as the last path segment, e.g. /catalog/PROD001
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) == 0 {
		http.Error(w, "missing product code", http.StatusBadRequest)
		return
	}
	code := parts[len(parts)-1]

	pd, err := h.service.GetProductDetails(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
