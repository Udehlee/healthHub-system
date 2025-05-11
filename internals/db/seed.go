package db

import (
	"log"
	"os"

	"github.com/Udehlee/healthHub-System/internals/models"
	"github.com/Udehlee/healthHub-System/utility"
)

func SeedData() models.User {
	hashedPassword, err := utility.HashPassword(os.Getenv("ADMIN_PASSWORD"))
	if err != nil {
		log.Fatalf("Failed to hash admin password: %v", err)
	}

	admin := models.User{
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: hashedPassword,
		Role:     os.Getenv("ADMIN_ROLE"),
	}
	return admin
}
