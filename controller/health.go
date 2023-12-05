package controller

import (
	"github.com/emersonfbarros/backend-challenge-klever/service"
	"github.com/gin-gonic/gin"
)

func Health(context *gin.Context) {
	health := service.Health()

	sendSuccess(context, *health)
}
