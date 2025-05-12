package db

import (
	"log"
	"os"

	"github.com/Udehlee/healthcare-Access/internals/models"
	"github.com/Udehlee/healthcare-Access/utility"
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
