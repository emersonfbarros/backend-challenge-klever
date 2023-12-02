package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize(port string) {
	// starts router with gin defaults
	router := gin.Default()

	// test route that will be removed in the future
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// runs server
	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
