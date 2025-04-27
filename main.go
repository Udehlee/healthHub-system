package main

import (
	"log"

	"github.com/Udehlee/healthHub-System/internals/api"
	"github.com/Udehlee/healthHub-System/internals/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	loadEnv()

	conn, err := db.InitDB()
	if err != nil {
		log.Fatal("error connecting to db")
	}

	admin := db.SeedData()
	conn.Save(&admin)

	h := api.NewHandler(conn)
	api.Routes(r, h)

	if err := r.Run(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	log.Println("successfully loaded .env file ")

}
