package api

import (
	"fmt"
	"net/http"

	"github.com/Udehlee/healthHub-System/internals/models"
	"github.com/Udehlee/healthHub-System/utility"
	"github.com/gin-gonic/gin"
)

// AdminAssign assigns a staff to a specific appointment.
func (h *Handler) AdminAssign(c *gin.Context) {
	appointmentID, err := utility.GetParamInt64(c, "id")
	if err != nil {
		fmt.Println("Appointment ID param:", appointmentID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment ID"})
		return
	}

	var assignReq models.AssignRequest
	if err := c.ShouldBindJSON(&assignReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind assign request"})
		return
	}

	adminID := c.GetInt64("user_id")
	if adminID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admin ID"})
		return
	}
	appointment := &models.Appointment{
		StaffID: assignReq.StaffID,
		// StaffRole:  assignReq.StaffRole,
		Status:     assignReq.Status,
		AssignedBy: &adminID,
	}

	err = h.Db.AssignStaff(appointmentID, appointment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to assign staff"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "staff successfully assigned",
		"appointments": appointment,
	})
}

// GetAllUsers retrieves all users
func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.Db.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get all users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully retrieved all users",
		"users":   users,
	})
}
