package service

import (
	"github.com/emersonfbarros/backend-challenge-klever/config"
	"github.com/emersonfbarros/backend-challenge-klever/model"
)

var logger *config.Logger

func InitService() {
	model.InitModel()
	logger = config.GetLogger("controller")
}
