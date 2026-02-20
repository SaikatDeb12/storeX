package main

import (
	"log"

	"github.com/SaikatDeb12/storeX/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env")
	}

	err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to database")
	}
}
