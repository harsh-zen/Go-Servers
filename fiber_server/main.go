package main

import (
	handler "fiber_server/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // Middleware for logging
    app.Use(logger.New())

    // Routes
    setupRoutes(app)

    log.Fatal(app.Listen(":8000"))
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
    api.Get("/appointments", handler.GetAppointments)
    api.Post("/appointments", handler.CreateAppointment)
    api.Put("/appointments/:id", handler.UpdateAppointment)
    api.Delete("/appointments/:id", handler.DeleteAppointment)
}

//requests:
//curl http://localhost:8003/api/v1/appointments
//curl -X POST -H "Content-Type: application/json" -d '{"title":"Meeting","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8003/api/v1/appointments
//curl -X PUT -H "Content-Type: application/json" -d '{"title":"Session on Go","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8003/api/v1/appointments/1
//curl -X DELETE http://localhost:8003/appointments/1
