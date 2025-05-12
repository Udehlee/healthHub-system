package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/Udehlee/healthcare-Access/internals/models"
	"github.com/golang-jwt/jwt"
)

var JWTKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateJWT generates token
func GenerateJWT(user *models.User) (string, error) {
	claims := models.Claims{
		ID:    user.UserID,
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(), // Token expires in 12 hours
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)

	}

	return tokenString, nil
}
