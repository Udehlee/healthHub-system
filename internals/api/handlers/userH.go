package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) BookAppointment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"error": "booking appointment jare!"})
}
