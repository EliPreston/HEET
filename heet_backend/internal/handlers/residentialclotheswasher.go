package handlers

import (
	"log"
	"net/http"

	"github.com/EliPreston/heet_backend/internal/services"

	"github.com/go-chi/render"
)

func ListResdientialClothesWashers(w http.ResponseWriter, r *http.Request) {
	residentialClothesWashers, err := services.GetResidentialClothesWashers()
	if err != nil {
		log.Printf("error querying residential_clothes_washers: %v", err)
		http.Error(w, "failed to fetch residential clothes washers", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, residentialClothesWashers) // returns JSON response
}
