package main

import (
	"log"

	"github.com/Udehlee/healthHub-System/internals/api"
	"github.com/Udehlee/healthHub-System/internals/db"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	conn, err := db.InitDB()
	if err != nil {
		log.Fatal("error connecting to db")
	}

	h := api.NewHandler(conn)

}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	log.Println("successfully .env file loaded ")

}
