package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	handler "chi_server/handlers"

	"github.com/go-chi/chi/v5"
)

func main(){
	router := chi.NewRouter()

	// Global prefix
	apiV1Prefix := "/api/v1"

	// Middleware
	router.Use(middleware.Logger)

	// Registering routes with global prefix
	router.Route(apiV1Prefix, func(r chi.Router) {
		r.Get("/appointments", handler.GetAppointments)
		r.Post("/appointments", handler.CreateAppointment)
		r.Put("/appointments/{id}", handler.UpdateAppointment)
		r.Delete("/appointments/{id}", handler.DeleteAppointment)
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Println("Server listening on", server.Addr)
	log.Fatal(server.ListenAndServe())

}

//Example Request
//curl http://localhost:8000/appointments
//curl -X POST -H "Content-Type: application/json" -d '{"title":"Meeting","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8000/appointments
//curl -X PUT -H "Content-Type: application/json" -d '{"title":"Session on Go","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8000/appointments/1
//curl -X DELETE http://localhost:8000/appointments/1

