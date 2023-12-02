package main

import (
	"log"
	"os"

	"github.com/emersonfbarros/backend-challenge-klever/config"
	"github.com/emersonfbarros/backend-challenge-klever/router"
)

func main() {
	// initialize router
	err := config.Init()

	if err != nil {
		log.Fatalf("Config initialization error: %v", err)
	}

	// get port from env
	port := os.Getenv("PORT")

	// initialize router
	router.Initialize(":" + port)
}
