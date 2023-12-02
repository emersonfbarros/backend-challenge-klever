package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Application is running!",
			})
		})
		v1.GET("/details/:address", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /details/:address",
			})
		})

		v1.GET("/balance/:address", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "GET /balance/:address",
			})
		})

		v1.POST("/send", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "POST /send",
			})
		})

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
