package controller

import (
	"github.com/emersonfbarros/backend-challenge-klever/config"
	"github.com/emersonfbarros/backend-challenge-klever/service"
)

var logger *config.Logger

func InitController() {
	service.InitService()
	logger = config.GetLogger("controller")
}
