package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AdminAssign(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"error": "assigning staff jare!"})
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
