package api

import (
	"log"
	"net/http"

	"github.com/Udehlee/healthcare-Access/internals/models"
	"github.com/gin-gonic/gin"
)

// BookAppointment books an appointment for the user
func (h *Handler) BookAppointment(c *gin.Context) {
	var AppointmentReq models.AppointmentRequest

	err := c.ShouldBindJSON(&AppointmentReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind appointment request"})
		return
	}

	Appointment := models.Appointment{
		PatientID: AppointmentReq.PatientID,
		Status:    "booked",
	}

	err = h.Db.SaveAppointment(&Appointment)
	if err != nil {
		log.Printf("Error saving appointment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to book appointment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "appointment successfully booked"})
}
