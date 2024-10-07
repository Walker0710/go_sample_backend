package config

import (
	"log"

	"github.com/joho/godotenv" // Import the godotenv package
)

// LoadConfig loads environment variables from a .env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
