package handler

import (
	model "fiber_server/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// In-memory storage for appointments (for demonstration purposes)
var appointments []model.Appointment

func init() {
	appointments = make([]model.Appointment, 1)
	appointments[0] = model.Appointment{
		ID:        1,
		Title:     "Meeting with Nemesh Sir",
		StartTime: time.Date(2024, 06, 10, 10, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 06, 10, 11, 0, 0, 0, time.UTC),
	}}

// GetAppointments handles the GET /appointments route
func GetAppointments(c *fiber.Ctx) error {
    return c.JSON(appointments)
}

// CreateAppointment handles the POST /appointments route
func CreateAppointment(c *fiber.Ctx) error {
    appointment := new(model.Appointment)

    if err := c.BodyParser(&appointment); err != nil {
        return err
    }

    appointment.ID = len(appointments) + 1
    appointments = append(appointments, *appointment)

    return c.JSON(appointment)
}

// UpdateAppointment handles the PUT /appointments/:id route
func UpdateAppointment(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id", -1) // Use -1 as the default value
    if err != nil || id == -1 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid appointment ID",
        })
    }

    var updatedAppointment model.Appointment
    if err := c.BodyParser(&updatedAppointment); err != nil {
        return err
    }

    for i, appointment := range appointments {
        if appointment.ID == id {
            updatedAppointment.ID = id
            appointments[i] = updatedAppointment
            break
        }
    }

    return c.JSON(updatedAppointment)
}

func DeleteAppointment(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id", -1) // Use -1 as the default value
    if err != nil || id == -1 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid appointment ID",
        })
    }

    for i, appointment := range appointments {
        if appointment.ID == id {
            appointments = append(appointments[:i], appointments[i+1:]...)
            break
        }
    }

    return c.SendStatus(fiber.StatusNoContent)
}