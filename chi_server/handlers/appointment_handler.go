package handler

import (
	model "chi_server/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

var appointments []model.Appointment

func init() {
	appointments = make([]model.Appointment, 1)
	appointments[0] = model.Appointment{
		ID:        1,
		Title:     "Meeting with Nemesh Sir",
		StartTime: time.Date(2024, 06, 10, 10, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 06, 10, 11, 0, 0, 0, time.UTC),
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