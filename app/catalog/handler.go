package catalog

import (
	"net/http"
	"strings"

	"github.com/mytheresa/go-hiring-challenge/app/api"
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
		api.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	products, total, err := h.service.GetProducts(req)
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"products": products,
		"total":    total,
	}

	api.OKResponse(w, response)
}

// HandleGetDetails returns full product details.
func (h *CatalogHandler) HandleGetDetails(w http.ResponseWriter, r *http.Request) {
	// Expect the code as the last path segment, e.g. /catalog/PROD001
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) == 0 {
		api.ErrorResponse(w, http.StatusBadRequest, "missing product code")
		return
	}
	code := parts[len(parts)-1]

	pd, err := h.service.GetProductDetails(code)
	if err != nil {
		api.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	api.OKResponse(w, pd)
}
