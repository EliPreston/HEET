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

		// Handle /appliances/dehumidifiers route
		r.Get("/dehumidifiers", handlers.ListDehumidifiers)
		r.Route("/dehumidifiers/{dehumidifier_id}", func(r chi.Router) {
			r.Get("/", handlers.GetDehumidifier)
		})
		// Nested route for single dehumidifier by ID
		// r.Route("/dehumidifiers/{id}", func(r chi.Router) {
		// 	r.Get("/", handlers.GetDehumidifierByID)
		// })

		// Handle /appliances/residentialclothesdryers route
		r.Get("/residentialclothesdryers", handlers.ListResidentialClothesDryers)
		// Nested route for single residentialclothesdryers by ID
		// r.Route("/residentialclothesdryers/{id}", func(r chi.Router) {
		// 	r.Get("/", handlers.GetApplianceByID)
		// })

		// Handle /appliances/residentialdishwashers route
		r.Get("/residentialdishwashers", handlers.ListResdientialDishwashers)
		// Nested route for single appliance by ID
		// r.Route("/residentialdishwashers/{id}", func(r chi.Router) {
		// 	r.Get("/", handlers.GetApplianceByID)
		// })

		// Handle /appliances/clotheswasherdryercombos route
		r.Get("/clotheswasherdryercombos", handlers.ListWasherDryerCombos)
		// Nested route for single clotheswasherdryercombo by ID
		// r.Route("/clotheswasherdryercombos/{id}", func(r chi.Router) {
		// 	r.Get("/", handlers.GetApplianceByID)
		// })

		// Handle /appliances/residentialclotheswashers route
		r.Get("/residentialclotheswashers", handlers.ListResdientialClothesWashers)
		// Nested route for single appliance by ID
		// r.Route("/residentialclotheswasher/{id}", func(r chi.Router) {
		// 	r.Get("/", handlers.GetApplianceByID)
		// })

		// Handle /appliances/categtoryofappliances route
		r.Get("/categtoryofappliances", func(w http.ResponseWriter, r *http.Request) { // r.Get("/", ListAppliances)
			w.Write([]byte("GET on 'categtoryofappliances' path: /appliances/categtoryofappliances"))
		})
		// Nested route for single appliance by ID
		// r.Route("/appliance/{id}", func(r chi.Router) {
		// 	r.Get("/", handlers.GetApplianceByID)
		// })

		// Handle /appliances/{id} route
		r.Route("/{id}", func(r chi.Router) {

			r.Get("/", func(w http.ResponseWriter, r *http.Request) { // r.Post("/", GetAppliance)
				w.Write([]byte("GET on 'appliances/{id}' path: /appliances/{id}"))
			})
			r.Put("/", func(w http.ResponseWriter, r *http.Request) { // r.Put("/", UpdateAppliance)
				w.Write([]byte("PUT on 'appliances/{id}' path: /appliances/{id}"))
			})
			r.Delete("/", func(w http.ResponseWriter, r *http.Request) { // r.Delete("/", DeleteAppliance)
				w.Write([]byte("DELETE on 'appliances/{id}' path: /appliances/{id}"))
			})
		})
	})

	// Start server
	fmt.Println("Server listening on port 8080 (http://localhost:8080)")
	http.ListenAndServe(":8080", r)

}
