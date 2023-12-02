package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize(port string) {
	// starts router with gin defaults
	router := gin.Default()

	initRoutes(router)

	// runs server
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
