package db

import (
	"os"

	"github.com/Udehlee/healthHub-System/internals/models"
)

func SeedData() models.User {
	admin := models.User{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Role:     os.Getenv("ADMIN_ROLE"),
	}

	return admin

}
