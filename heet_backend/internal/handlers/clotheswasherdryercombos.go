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

func ListWasherDryerCombos(w http.ResponseWriter, r *http.Request) {
	washerDryerCombos, err := services.GetWasherDryerCombos()
	if err != nil {
		log.Printf("error querying washer_dryer_combos: %v", err)
		http.Error(w, "failed to fetch washer-dryer combos", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, washerDryerCombos) // returns JSON response
}

func GetClothesWasherDryerCombo(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "clotheswasherdryercombo_id")
	id_float, err := strconv.ParseFloat(id, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}

	washerdryer, err := services.GetWasherDryerComboByID(id_float)

	if err != nil {
		log.Printf("error querying clothes washer-dryer combo with id %s: %v", id, err)
		http.Error(w, "failed to fetch clothes washer-dryer combo", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, washerdryer) // returns JSON response
}
