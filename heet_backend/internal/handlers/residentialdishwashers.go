package handlers

import (
	"log"
	"net/http"

	"github.com/EliPreston/heet_backend/internal/services"

	"github.com/go-chi/render"
)

func ListResdientialDishwashers(w http.ResponseWriter, r *http.Request) {
	residentialDishwashers, err := services.GetResidentialDishwashers()
	if err != nil {
		log.Printf("error querying residential_dishwashers: %v", err)
		http.Error(w, "failed to fetch residential dishwashers", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, residentialDishwashers) // returns JSON response
}
