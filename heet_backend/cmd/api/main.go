package main

import (
	"fmt"
	"net/http"

	// For Postgres connection
	"github.com/EliPreston/heet_backend/internal/db"
	"github.com/EliPreston/heet_backend/internal/handlers"

	// For routing requests
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {

	// Connect to Postgres DB, app will exit if no connection can be made
	db.ConnectToDB()

	r := chi.NewRouter()

	// Apply middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Define routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root path: /"))
	})

	r.Route("/appliances", func(r chi.Router) {

		// Handle /appliances route
		r.Get("/", func(w http.ResponseWriter, r *http.Request) { // r.Get("/", ListAppliances)
			w.Write([]byte("GET on 'appliances' path: /appliances"))
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) { // r.Post("/", CreateAppliance)
			w.Write([]byte("POST on 'appliances' path: /appliances"))
		})

		// Handle dehumidifiers routes
		r.Get("/dehumidifiers", handlers.ListDehumidifiers)
		r.Route("/dehumidifiers/{dehumidifier_id}", func(r chi.Router) {
			r.Get("/", handlers.GetDehumidifier)
		})

		// Handle residential clothes dryers routes
		r.Get("/residentialclothesdryers", handlers.ListResidentialClothesDryers)
		r.Route("/residentialclothesdryers/{residentialclothesdryer_id}", func(r chi.Router) {
			r.Get("/", handlers.GetResidentialClothesDryer)
		})

		// Handle dishwashers routes
		r.Get("/residentialdishwashers", handlers.ListResdientialDishwashers)
		r.Route("/residentialdishwashers/{residentialdishwasher_id}", func(r chi.Router) {
			r.Get("/", handlers.GetResidentialDishwasher)
		})

		// Handle clothes washer-dryer combos routes
		r.Get("/clotheswasherdryercombos", handlers.ListWasherDryerCombos)
		r.Route("/clotheswasherdryercombos/{clotheswasherdryercombo_id}", func(r chi.Router) {
			r.Get("/", handlers.GetClothesWasherDryerCombo)
		})

		// Handle residential clothes washers routes
		r.Get("/residentialclotheswashers", handlers.ListResdientialClothesWashers)
		r.Route("/residentialclotheswashers/{residentialclotheswasher_id}", func(r chi.Router) {
			r.Get("/", handlers.GetResidentialClothesWasher)
		})

		//------------------------------------------
		//------------------------------------------
		//------------------------------------------
		// // Handle /appliances/categtoryofappliances route
		// r.Get("/categtoryofappliances", func(w http.ResponseWriter, r *http.Request) { // r.Get("/", ListAppliances)
		// 	w.Write([]byte("GET on 'categtoryofappliances' path: /appliances/categtoryofappliances"))
		// })
		// // Nested route for single appliance by ID
		// // r.Route("/appliance/{id}", func(r chi.Router) {
		// // 	r.Get("/", handlers.GetApplianceByID)
		// // })
		// // Handle /appliances/{id} route
		// r.Route("/{id}", func(r chi.Router) {

		// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) { // r.Post("/", GetAppliance)
		// 		w.Write([]byte("GET on 'appliances/{id}' path: /appliances/{id}"))
		// 	})
		// 	r.Put("/", func(w http.ResponseWriter, r *http.Request) { // r.Put("/", UpdateAppliance)
		// 		w.Write([]byte("PUT on 'appliances/{id}' path: /appliances/{id}"))
		// 	})
		// 	r.Delete("/", func(w http.ResponseWriter, r *http.Request) { // r.Delete("/", DeleteAppliance)
		// 		w.Write([]byte("DELETE on 'appliances/{id}' path: /appliances/{id}"))
		// 	})
		// })
	})

	// Start server
	fmt.Println("Server listening on port 8080 (http://localhost:8080)")
	http.ListenAndServe(":8080", r)

}
