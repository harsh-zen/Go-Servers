package handler

import (
	"encoding/json"
	"fmt"
	"log"
	model "native_server/models"
	"net/http"
	"strconv"
	"time"
)

var appointments []model.Appointment

func init() {
	log.Println("Initializing 100000 appointments...")
    appointments = make([]model.Appointment, 100000) 
    for i := 0; i < 100000; i++ { 
        appointments[i] = model.Appointment{
            ID:        i+1,
            Title:     fmt.Sprintf("Meeting %d", i+1),
            StartTime: time.Now(),
            EndTime:   time.Now().Add(time.Hour),
        }
    }
}

//  GET /appointments
func GetAppointments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(appointments)
}

// POST /appointments
func CreateAppointment(w http.ResponseWriter, r *http.Request){
	var appointment model.Appointment
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	appointment.ID = len(appointments) + 1
	appointments = append(appointments, appointment)
	json.NewEncoder(w).Encode(appointment)
}

// PUT /appointments/{id}
func UpdateAppointment(w http.ResponseWriter, r *http.Request){
	var updatedAppointment model.Appointment
	err := json.NewDecoder(r.Body).Decode(&updatedAppointment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	recieved_id, err := strconv.Atoi(r.PathValue("id"))
	for i, appointment := range appointments {
		if appointment.ID == recieved_id {
			updatedAppointment.ID = recieved_id
			appointments[i] = updatedAppointment
			json.NewEncoder(w).Encode(updatedAppointment)
			return
		}
	}
	http.Error(w, "Appointment not found", http.StatusNotFound)
}

// DELETE /appointments/{id}
func DeleteAppointment(w http.ResponseWriter, r *http.Request){
	recieved_id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, appointment := range appointments {
		if appointment.ID == recieved_id {
			appointments = append(appointments[:i], appointments[i+1:]...)
			json.NewEncoder(w).Encode(appointments)
			return
		}
	}
	http.Error(w, "Appointment not found", http.StatusNotFound)
}
