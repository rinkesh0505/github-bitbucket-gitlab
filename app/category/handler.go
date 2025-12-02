package category

import (
	"encoding/json"
	"net/http"
)

type CategoryHandler struct {
	service *CategoryService
}

func NewCategoryHandler(service *CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// HandleGetAll handles GET /categories requests.
func (h *CategoryHandler) HandleGetAll(w http.ResponseWriter, r *http.Request) {
	cats, err := h.service.GetAllCategories()
	if err != nil {
		http.Error(w, "failed to fetch categories", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cats); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleCreate handles POST /categories requests to create a new category.
func (h *CategoryHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var cat Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := cat.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.CreateCategory(cat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(cat); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
