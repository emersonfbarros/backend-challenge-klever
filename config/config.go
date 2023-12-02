package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Init() error {
	err := godotenv.Load()

	if err != nil {
		return fmt.Errorf("Error loading .env file. Error: %v", err)
	}

	return nil
}
