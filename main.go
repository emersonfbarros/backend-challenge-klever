package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	// chanel to indicate when http server is closed
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// initialize server
	srv := router.Initialize(":" + port)

	// waits until receiving an interrupt signal
	<-done

	// close the HTTP server with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Errorf("Server Shutdown Failed:%+v", err)
	}
	logger.Infof("Server Exited Properly")
}
