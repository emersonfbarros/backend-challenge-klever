package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Init() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file. Error: %v", err)
	}

	return nil
}
