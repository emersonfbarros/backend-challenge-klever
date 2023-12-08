package router

import (
	"net/http"
	"os"

	"github.com/emersonfbarros/backend-challenge-klever/config"
	"github.com/gin-gonic/gin"
)

var logger config.ILogger

func Initialize(port string) *http.Server {
	logger = config.GetLogger("router")
	// starts router with gin defaults
	router := gin.Default()

	initRoutes(router)

	// Create the server
	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// Run the server in a goroutine so it doesn't block
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("Failed to run server: %v", err)
			os.Exit(1)
		}
	}()

	logger.Infof("Server running on port %s", port)

	return srv
}
