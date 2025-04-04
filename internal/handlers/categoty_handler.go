package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	services "github.com/breadwithmeth/naliv_go/internal/services"
)

type CategoryHandler struct {
	service *services.CategoryService
}
func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	if service == nil {
		panic("category service is nil")
	}
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	businessIdStr := r.URL.Query().Get("business_id")
	businessId, err := strconv.Atoi(businessIdStr)
	if err != nil {
		http.Error(w, "Invalid business ID", http.StatusBadRequest)
		return
	}

	categories, err := h.service.GetCategories(businessId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}