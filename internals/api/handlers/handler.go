package api

import (
	"net/http"

	"github.com/Udehlee/healthHub-System/internals/db"
	"github.com/Udehlee/healthHub-System/internals/models"
	"github.com/Udehlee/healthHub-System/utility"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Db db.Store
}

func NewHandler(db db.Store) *Handler {
	return &Handler{
		Db: db,
	}
}

// Register creates a new user account
func (h *Handler) Register(c *gin.Context) {
	var RegisterReq models.User

	err := c.ShouldBindJSON(&RegisterReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind register request"})
		return
	}

	Email, err := h.Db.CheckEmail(RegisterReq.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB error"})
	}
	if Email != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already in use"})
		return
	}

	hashedpwd, err := utility.HashPassword(RegisterReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	defaultRole := &models.Role{RoleName: "patient"}
	newUser := models.User{
		FirstName: RegisterReq.FirstName,
		LastName:  RegisterReq.LastName,
		Email:     RegisterReq.Email,
		Password:  hashedpwd,
		Role:      defaultRole, // Assign the Role instance here
		Gender:    RegisterReq.Gender,
		Address:   RegisterReq.Address,
	}

	err = h.Db.Save(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully saved user"})
}

// Login logs in an existing user
func (h *Handler) Login(c *gin.Context) {
	var LoginReq models.LoginRequest

	err := c.ShouldBindJSON(&LoginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind login request"})
		return
	}

	user, err := h.Db.CheckEmail(LoginReq.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found, create an account"})
		return
	}

	err = utility.ComparePasswordHash(user.Password, LoginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password"})
	}

	c.JSON(http.StatusOK, gin.H{"message": " logged in sucessefully"})

}
