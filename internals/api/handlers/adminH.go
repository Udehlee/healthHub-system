package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AdminAssign(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"error": "assigning staff jare!"})
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"error": "getting all users jare!"})
}
