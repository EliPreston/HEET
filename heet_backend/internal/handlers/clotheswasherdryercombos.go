package handlers

import (
	"log"
	"net/http"

	"github.com/EliPreston/heet_backend/internal/services"

	"github.com/go-chi/render"
)

func ListWasherDryerCombos(w http.ResponseWriter, r *http.Request) {
	washerDryerCombos, err := services.GetWasherDryerCombos()
	if err != nil {
		log.Printf("error querying washer_dryer_combos: %v", err)
		http.Error(w, "failed to fetch washer-dryer combos", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, washerDryerCombos) // returns JSON response
}
