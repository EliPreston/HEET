package main

import (
	"fmt"
	"net/http"

	// For Postgres connection
	"backend/services"

	// For routing requests
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {

	// Connect to Postgres DB, app will exit if no connection can be made
	services.ConnectToDB()

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
