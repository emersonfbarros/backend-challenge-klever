package controller

import (
	"github.com/emersonfbarros/backend-challenge-klever/config"
	"github.com/emersonfbarros/backend-challenge-klever/model"
	"github.com/emersonfbarros/backend-challenge-klever/service"
)

var logger *config.Logger
var services service.IServices
var models model.IModels
var fetcher model.IFetcher

func InitController() {
	service.InitService()
	models = model.NewModels()
	fetcher = model.NewFetcher()
	services = service.NewServices()
	logger = config.GetLogger("controller")
}
