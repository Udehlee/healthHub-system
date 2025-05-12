package routes

import (
	api "github.com/Udehlee/healthcare-Access/internals/api/handlers"
	"github.com/Udehlee/healthcare-Access/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, h *api.Handler) {
	r.GET("/", h.Index)
	r.POST("/api/register", h.Register)
	r.POST("/api/login", h.Login)

	//admin routes
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware(), middleware.Role("admin"))
	{
		adminRoutes.PATCH("/appointments/:id", h.AdminAssign)
		adminRoutes.GET("/users", h.GetAllUsers)
	}

	//deafult user(patient)
	userRoutes := r.Group("/user")
	userRoutes.Use(middleware.AuthMiddleware(), middleware.Role("patient"))
	{
		userRoutes.POST("/appointments/book-appointment", h.BookAppointment)
	}

	staffRoutes := r.Group("/staff")
	staffRoutes.Use(middleware.AuthMiddleware(), middleware.Role("staff"))
	{
		staffRoutes.GET("/appointments/assigned", h.ViewAssigned)
	}
}
