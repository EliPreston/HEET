package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/EliPreston/heet_backend/internal/services"

	"github.com/go-chi/chi/v5"
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

func GetResidentialDishwasher(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "residentialdishwasher_id")
	id_float, err := strconv.ParseFloat(id, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}

	dishwasher, err := services.GetDishwasherByID(id_float)

	if err != nil {
		log.Printf("error querying residential dishwasher with id %s: %v", id, err)
		http.Error(w, "failed to fetch residential dishwasher dryer", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, dishwasher) // returns JSON response
}
