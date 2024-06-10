package main

import (
	handler "iris_server/handlers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
)

func main(){
	app := iris.New()
	app.Use(iris.Compression)
	app.Use(logger.New())


	apiV1 := app.Party("/api/v1")
	{
		apiV1.Get("/appointments", handler.GetAppointments)
		apiV1.Post("/appointments", handler.CreateAppointment)
		apiV1.Put("/appointments/{id}", handler.UpdateAppointment)
		apiV1.Delete("/appointments/{id}", handler.DeleteAppointment)
	}

	app.Listen(":8000")
}

//Example Request
//curl http://localhost:8000/api/v1/appointments
//curl -X POST -H "Content-Type: application/json" -d '{"title":"Meeting","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8000/api/v1/appointments
//curl -X PUT -H "Content-Type: application/json" -d '{"title":"Session on Go","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8000/api/v1/appointments/1
//curl -X DELETE http://localhost:8000/api/v1/appointments/1

