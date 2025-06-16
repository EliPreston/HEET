package handlers

import (
	"log"
	"net/http"

	"backend/services" // where your DB connection lives

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
