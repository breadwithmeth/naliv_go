package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/breadwithmeth/naliv_go/models"
)

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	items := []models.Item{
		{ID: 1, Name: "Item 1", Price: 100},
		{ID: 2, Name: "Item 2", Price: 200},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
