package main

import (
	"log"

	api "github.com/Udehlee/healthcare-Access/internals/api/handlers"
	"github.com/Udehlee/healthcare-Access/internals/api/routes"
	"github.com/Udehlee/healthcare-Access/internals/db"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	conn, err := db.InitDB()
	if err != nil {
		log.Fatal("error connecting to db")
	}

	admin := db.SeedData()
	conn.Save(&admin)

	h := api.NewHandler(conn)
	routes.Routes(r, h)

	if err := r.Run(":8000"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}

// func loadEnv() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("failed to load .env file")
// 	}

// 	log.Println("successfully loaded .env file ")

// }
