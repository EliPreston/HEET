package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/EliPreston/heet_backend/internal/services" // where your DB connection lives

	"github.com/go-chi/chi/v5"
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

func GetResidentialClothesDryer(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "residentialclothesdryer_id")
	id_float, err := strconv.ParseFloat(id, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}

	dryer, err := services.GetResidentialClothesDryerByID(id_float)

	if err != nil {
		log.Printf("error querying residential clothes dryer with id %s: %v", id, err)
		http.Error(w, "failed to fetch residential clothes dryer", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, dryer) // returns JSON response
}
