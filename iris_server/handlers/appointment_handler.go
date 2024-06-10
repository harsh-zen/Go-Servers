package handler

import (
	model "iris_server/models"
	"net/http"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
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

func GetAppointments(ctx iris.Context) {
	ctx.JSON(appointments)
}

// CreateAppointment handles POST /api/v1/appointments
func CreateAppointment(ctx iris.Context) {
	var appointment model.Appointment
	if err := ctx.ReadJSON(&appointment); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	appointment.ID = len(appointments) + 1
	appointments = append(appointments, appointment)
	ctx.StatusCode(http.StatusCreated)
	ctx.JSON(appointment)
}

// UpdateAppointment handles PUT /api/v1/appointments/{id}
func UpdateAppointment(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid appointment ID"})
		return
	}

	var updatedAppointment model.Appointment
	if err := ctx.ReadJSON(&updatedAppointment); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	for i, appointment := range appointments {
		if appointment.ID == id {
			updatedAppointment.ID = id
			appointments[i] = updatedAppointment
			ctx.JSON(updatedAppointment)
			return
		}
	}

	ctx.StatusCode(http.StatusNotFound)
	ctx.JSON(iris.Map{"error": "Appointment not found"})
}

// DeleteAppointment handles DELETE /api/v1/appointments/{id}
func DeleteAppointment(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid appointment ID"})
		return
	}

	for i, appointment := range appointments {
		if appointment.ID == id {
			appointments = append(appointments[:i], appointments[i+1:]...)
			ctx.StatusCode(http.StatusNoContent)
			return
		}
	}

	ctx.StatusCode(http.StatusNotFound)
	ctx.JSON(iris.Map{"error": "Appointment not found"})
}
