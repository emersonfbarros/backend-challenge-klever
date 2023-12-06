package model

import (
	"os"

	"github.com/emersonfbarros/backend-challenge-klever/config"
)


var logger *config.Logger

func InitModel() {
	logger = config.GetLogger("model")
}
