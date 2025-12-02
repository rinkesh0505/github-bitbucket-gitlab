package category

import (
	"encoding/json"
	"net/http"

	"github.com/mytheresa/go-hiring-challenge/app/api"
)

type CategoryHandler struct {
	service CategoryService
}

func NewCategoryHandler(service CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleGetAll handles GET /categories requests.
func (h *CategoryHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	cats, err := h.service.GetAllCategories()
	if err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, "failed to fetch categories")
		return
	}
	api.OKResponse(w, cats)
}

// HandleCreate handles POST /categories requests to create a new category.
func (h *CategoryHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var cat Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		api.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if err := cat.Validate(); err != nil {
		api.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateCategory(cat); err != nil {
		api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	api.OKResponse(w, cat)
}
