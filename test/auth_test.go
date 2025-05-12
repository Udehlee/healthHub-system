package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Udehlee/healthcare-Access/internals/models"
	"github.com/Udehlee/healthcare-Access/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	user := &models.User{
		UserID: 1,
		Email:  "pita@email.com",
		Role:   "admin",
	}

	tokenStr, err := middleware.GenerateJWT(user)
	assert.NoError(t, err, "should not return error generating token")
	assert.NotEmpty(t, tokenStr, "token string should not be empty")

	token, err := jwt.ParseWithClaims(tokenStr, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return middleware.JWTKey, nil
	})

	assert.NoError(t, err, "should successfully parse token")
	assert.True(t, token.Valid, "token should be valid")

	claims, ok := token.Claims.(*models.Claims)
	assert.True(t, ok, "claims should be of type *models.Claims")

	assert.Equal(t, user.UserID, claims.ID)
	assert.Equal(t, user.Email, claims.Email)
	assert.Equal(t, user.Role, claims.Role)
	assert.WithinDuration(t, time.Now().Add(12*time.Hour), time.Unix(claims.ExpiresAt, 0), time.Minute)
}

func TestRole(t *testing.T) {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("claims", &models.Claims{
			Email: "user@example.com",
			Role:  "user",
		})
		c.Next()
	})

	r.GET("/admin", middleware.Role("admin"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome, admin!"})
	})

	req, _ := http.NewRequest("GET", "/admin", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusForbidden, resp.Code)
	assert.Contains(t, resp.Body.String(), "Forbidden")

}
