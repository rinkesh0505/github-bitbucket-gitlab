package catalog

import (
	"encoding/json"
	"net/http"
)

type CatalogHandler struct {
	service *CatalogService
}

func NewCatalogHandler(service *CatalogService) *CatalogHandler {
	return &CatalogHandler{
		service: service,
	}
}

func (h *CatalogHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		Products: products,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
