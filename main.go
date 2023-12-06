package main

import (
	"os"

	"github.com/emersonfbarros/backend-challenge-klever/config"
	"github.com/emersonfbarros/backend-challenge-klever/router"
)

var logger config.ILogger

func main() {
	logger = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		os.Exit(1)
	}

	// get port from env
	port := os.Getenv("PORT")

	// initialize router
	router.Initialize(":" + port)
}
