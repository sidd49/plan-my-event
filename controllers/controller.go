package controllers

import "example.com/plan-my-event/service"

type Controller interface {
	IEventsController
	IRegisterController
	IUsersController
}

type controllerImpl struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return controllerImpl{
		service: service,
	}
}
