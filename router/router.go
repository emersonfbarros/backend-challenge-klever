package router

import (
	"os"

	"github.com/emersonfbarros/backend-challenge-klever/config"
	"github.com/gin-gonic/gin"
)

var logger *config.Logger

func Initialize(port string) {
	logger = config.GetLogger("router")
	// starts router with gin defaults
	router := gin.Default()

	initRoutes(router)

	// runs server
	logger.Infof("Server running on port %s", port)
	if err := router.Run(port); err != nil {
		logger.Errorf("Failed to run server: %v", err)
		os.Exit(1)
	}
}
