package api

import (
	api "github.com/Udehlee/healthHub-System/internals/api/handlers"
	"github.com/Udehlee/healthHub-System/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, h *api.Handler) {

	r.POST("/api/register", h.Register)
	r.POST("/api/login", h.Login)

	//admin routes
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
	{
		adminRoutes.POST("/assign-patient-to-staff", h.AdminAssign)
		adminRoutes.GET("/Get-all-users", h.GetAllUsers)
	}

	//deafult user(patient)
	userRoutes := r.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("patient"))
	{
		userRoutes.POST("/book-appointment", h.BookAppointment)
	}
}
