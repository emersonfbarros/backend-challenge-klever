package model

import "github.com/emersonfbarros/backend-challenge-klever/config"

var logger *config.Logger

func InitModel() {
	logger = config.GetLogger("controller")
}
