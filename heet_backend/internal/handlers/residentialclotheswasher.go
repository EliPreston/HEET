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

func ListResdientialClothesWashers(w http.ResponseWriter, r *http.Request) {
	residentialClothesWashers, err := services.GetResidentialClothesWashers()
	if err != nil {
		log.Printf("error querying residential_clothes_washers: %v", err)
		http.Error(w, "failed to fetch residential clothes washers", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, residentialClothesWashers) // returns JSON response
}

func GetResidentialClothesWasher(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "residentialclotheswasher_id")
	id_float, err := strconv.ParseFloat(id, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}

	clotheswasher, err := services.GetResidentialClothesWasherByID(id_float)

	if err != nil {
		log.Printf("error querying residential clothes washer with id %s: %v", id, err)
		http.Error(w, "failed to fetch residential clothes washer", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, clotheswasher) // returns JSON response
}
