package config

import (
	"time"

	"github.com/joho/godotenv"
)

var AppStartTime time.Time

func Init() error {
	AppStartTime = time.Now()

	err := godotenv.Load()

	if err != nil {
		return err
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
