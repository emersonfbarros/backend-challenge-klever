package model

import "github.com/emersonfbarros/backend-challenge-klever/config"

var logger config.ILogger

func InitModel() {
	logger = config.GetLogger("controller")
}
