package handler

import (
	model "gin_server/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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
func GetAppointments(c *gin.Context) {
	c.JSON(http.StatusOK, appointments)
}

// POST /appointments
func CreateAppointment(c *gin.Context){
	var appointment model.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment.ID = len(appointments) + 1
	appointments = append(appointments, appointment)
	c.JSON(http.StatusCreated, appointment)
}

// PUT /appointments/{id}
func UpdateAppointment(c *gin.Context){
	idParam := c.Param("id")
	recieved_id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}
	var updatedAppointment model.Appointment
	if err := c.ShouldBindJSON(&updatedAppointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, appointment := range appointments {
		if appointment.ID == recieved_id {
			updatedAppointment.ID = recieved_id
			appointments[i] = updatedAppointment
			c.JSON(http.StatusOK, updatedAppointment)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
}

// DELETE /appointments/{id}
func DeleteAppointment(c *gin.Context) {
	idParam := c.Param("id")
	recieved_id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID"})
		return
	}

	for i, appointment := range appointments {
		if appointment.ID == recieved_id {
			appointments = append(appointments[:i], appointments[i+1:]...)
			c.JSON(http.StatusOK, appointments)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
}