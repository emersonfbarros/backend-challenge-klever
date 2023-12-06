package controller

import (
	"github.com/gin-gonic/gin"
)

func Health(context *gin.Context) {
	health := services.Health(fetcher)

	sendSuccess(context, *health)
}
