package service

import (
	"example.com/plan-my-event/repository"
)

type Service interface {
	IEventsService
	IRegisterService
	IUserService
}

type serviceImpl struct {
	Store repository.Repository
}

func NewService(store repository.Repository) Service {
	return serviceImpl{
		Store: store,
	}
}
