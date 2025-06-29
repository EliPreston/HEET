package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	// where DB connection lives
	"github.com/EliPreston/heet_backend/internal/services"

	// For routing requests
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ListDehumidifiers(w http.ResponseWriter, r *http.Request) {
	dehumidifiers, err := services.GetDehumidifiers()
	if err != nil {
		log.Printf("error querying dehumidifiers: %v", err)
		http.Error(w, "failed to fetch dehumidifiers", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, dehumidifiers) // returns JSON response
}

func GetDehumidifier(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "dehumidifier_id")
	id_float, err := strconv.ParseFloat(id, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}

	dehumidifier, err := services.GetDehumidifierByID(id_float)

	if err != nil {
		log.Printf("error querying dehumidifier with id %s: %v", id, err)
		http.Error(w, "failed to fetch dehumidifier", http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, dehumidifier) // returns JSON response
}
