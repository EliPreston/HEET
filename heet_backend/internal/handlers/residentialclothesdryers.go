package handlers

import (
	"log"
	"net/http"

	"github.com/EliPreston/heet_backend/internal/services" // where your DB connection lives

	"github.com/go-chi/render"
)

func ListResidentialClothesDryers(w http.ResponseWriter, r *http.Request) {
	residentialClothesDryers, err := services.GetResidentialClothesDryers()
	if err != nil {
		log.Printf("error querying residentialclothesdryers: %v", err)
		http.Error(w, "failed to fetch residential clothes dryers", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, residentialClothesDryers) // returns JSON response
}
