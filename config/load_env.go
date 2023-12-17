package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
