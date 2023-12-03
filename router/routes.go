package router

import (
	"net/http"

	"github.com/emersonfbarros/backend-challenge-klever/controller"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	controller.InitController()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Application is running!",
			})
		})

		v1.GET("/details/:address", controller.Details)

		v1.GET("/balance/:address", controller.Balance)

		v1.POST("/send", controller.Send)

		v1.GET("/tx/:tx", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /tx/:tx",
			})
		})

		v1.GET("/health", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /health",
			})
		})
	}
}
