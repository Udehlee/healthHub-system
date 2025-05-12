package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ViewAssigned views all the assigned appointments
func (h *Handler) ViewAssigned(c *gin.Context) {
	appointments, err := h.Db.GetAssignedAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch assigned appointments"})
		return
	}

	c.JSON(http.StatusOK, appointments)
}
