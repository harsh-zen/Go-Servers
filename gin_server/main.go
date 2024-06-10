package main

import (
	handler "gin_server/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){

    router := gin.Default()
	router.Use(gin.Logger())

	apiV1 := router.Group("/api/v1")

	{
		apiV1.GET("/appointments", handler.GetAppointments)
		apiV1.POST("/appointments", handler.CreateAppointment)
		apiV1.PUT("/appointments/:id", handler.UpdateAppointment)
		apiV1.DELETE("/appointments/:id", handler.DeleteAppointment)
	}
	server := &http.Server{
		Addr:    ":8000",
		Handler: router,

	}
	log.Println("Server listening on" + server.Addr)
	server.ListenAndServe()
}

//Example Request
//curl http://localhost:8000/api/v1/appointments
//curl -X POST -H "Content-Type: application/json" -d '{"title":"Meeting","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8000/api/v1/appointments
//curl -X PUT -H "Content-Type: application/json" -d '{"title":"Session on Go","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8000/api/v1/appointments/1
//curl -X DELETE http://localhost:8000/api/v1/appointments/1

