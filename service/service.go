package service

import "github.com/emersonfbarros/backend-challenge-klever/config"

var logger *config.Logger

func InitService() {
	logger = config.GetLogger("controller")
}
