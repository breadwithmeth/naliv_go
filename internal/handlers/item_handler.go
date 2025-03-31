package handlers

import (
	"encoding/json"
	"net/http"

	// "github.com/breadwithmeth/naliv_go/internal/models"
	services "github.com/breadwithmeth/naliv_go/internal/services"
)

type ItemHandler struct {
	service *services.ItemService
}

func NewItemHandler(service *services.ItemService) *ItemHandler {
	if service == nil {
		panic("item service is nil")
	}
	return &ItemHandler{
		service: service,
	}
}

type ItemFilter struct {
	BusinessID int `json:"business_id"`
	CategoryID int `json:"category_id"`
}

func (h *ItemHandler) GetAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	var filter ItemFilter
	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	items, err := h.service.GetItems(filter.BusinessID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
