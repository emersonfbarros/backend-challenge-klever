package controller

import "github.com/emersonfbarros/backend-challenge-klever/config"

var logger *config.Logger

func InitController() {
	logger = config.GetLogger("controller")
}
