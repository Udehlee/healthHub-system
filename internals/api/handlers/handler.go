package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Udehlee/healthcare-Access/internals/db"
	"github.com/Udehlee/healthcare-Access/internals/models"
	"github.com/Udehlee/healthcare-Access/middleware"
	"github.com/Udehlee/healthcare-Access/utility"
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

func (h *Handler) Index(c *gin.Context) {
	c.String(200, "Welcome Home")
}

// Register creates a new user account
func (h *Handler) Register(c *gin.Context) {
	var RegisterReq models.User

	if err := c.ShouldBindJSON(&RegisterReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind register request"})
		return
	}

	hashedpwd, err := utility.HashPassword(RegisterReq.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	newUser := models.User{
		FirstName: RegisterReq.FirstName,
		LastName:  RegisterReq.LastName,
		Email:     RegisterReq.Email,
		Password:  hashedpwd,
		Role:      "",
		Gender:    RegisterReq.Gender,
		Address:   RegisterReq.Address,
	}

	if err := h.Db.Save(&newUser); err != nil {
		log.Printf("saving error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user"})
		return
	}

	token, err := middleware.GenerateJWT(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
		"token":   token,
		"user": gin.H{
			"id":        newUser.UserID,
			"firstname": newUser.FirstName,
			"lastname":  newUser.LastName,
			"email":     newUser.Email,
			"gender":    newUser.Gender,
			"address":   newUser.Address,
			"role":      newUser.Role,
		},
	})
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
		fmt.Printf("Error in ComparePasswordHash: %v\n", err) // Logging the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password"})
		return
	}

	accesstoken, err := middleware.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logged in successfully",
		"token":   accesstoken,
		"user": gin.H{
			"id":        user.UserID,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
			"email":     user.Email,
			"role":      user.Role,
		},
	})
}
