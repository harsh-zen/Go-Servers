package main

import (
	"log"
	handler "native_server/handlers"
	middleware "native_server/utils"
	"net"
	"net/http"
	"os"
)

func main(){

	router := http.NewServeMux()

	//Trying to implement group based routing
	//Version 1 (JSON)
	apiV1Prefix := "/api/v1"
	apiV1Router := http.NewServeMux()

	apiV1Router.HandleFunc("GET /appointments", handler.GetAppointments)

	apiV1Router.HandleFunc("POST /appointments", handler.CreateAppointment)

	apiV1Router.HandleFunc("PUT /appointments/{id}", handler.UpdateAppointment)

	apiV1Router.HandleFunc("DELETE /appointments/{id}", handler.DeleteAppointment)

	router.Handle(apiV1Prefix+"/", http.StripPrefix(apiV1Prefix, apiV1Router))

	//Version 2 (Protobuf)
	apiV2Prefix := "/api/v2"
	apiV2Router := http.NewServeMux()
	apiV2Router.HandleFunc("GET /appointments", handler.GetAppointmentsProto)
	router.Handle(apiV2Prefix+"/", http.StripPrefix(apiV2Prefix, apiV2Router))

	port := ":8000"
	if !isPortAvailable(port) {
		log.Fatalf("Port %s is already in use", port)
		os.Exit(1)
	}


	server := &http.Server{
		Addr:    ":8000",
		Handler: middleware.Logging(router),

	}
	log.Println("Server listening on" + server.Addr)
	server.ListenAndServe()
}
func isPortAvailable(port string) bool {
	listener, err := net.Listen("tcp", port)
	if err!= nil {
		return false
	}
	defer listener.Close()
	return true
}

//Example Request
//curl http://localhost:8000/api/v1/appointments
//curl -X POST -H "Content-Type: application/json" -d '{"title":"Meeting","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8000/api/v1/appointments
//curl -X PUT -H "Content-Type: application/json" -d '{"title":"Session on Go","start_time":"2024-06-01T09:00:00Z","end_time":"2024-06-01T10:00:00Z"}' http://localhost:8000/api/v1/appointments/1
//curl -X DELETE http://localhost:8000/api/v1/appointments/1
//curl http://localhost:8000/api/v2/appointments

