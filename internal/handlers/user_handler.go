package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	services "github.com/breadwithmeth/naliv_go/internal/services"
	"github.com/breadwithmeth/naliv_go/internal/transport/rest/middleware"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	if service == nil {
		panic("user service is nil")
	}
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserAddressesHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token required", http.StatusUnauthorized)
		return
	}

	// id, err := h.service.GetUserByToken(token)
	id, ok := middleware.GetUserIDFromContext(r.Context())

	if !ok {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	println("User ID from context:", id)
	addresses, err := h.service.GetUserAddresses(id)
	if err != nil {
		http.Error(w, "Failed to get user addresses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addresses)
}
