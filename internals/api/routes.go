package api

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine, h *Handler) {

	r.POST("/api/register", h.Register)
	r.POST("/api/login", h.Login)

}
