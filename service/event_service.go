package service

import (
	"errors"
	"log"

	"example.com/plan-my-event/dto"
	models "example.com/plan-my-event/models"
	"example.com/plan-my-event/utils"
	"github.com/gin-gonic/gin"
)

type IEventsService interface {
	GetEvents(*gin.Context) ([]dto.EventDto, error)
	GetEvent(*gin.Context, string) (*models.Event, *utils.CustomError)
	CreateEvents(*gin.Context, []models.Event) *utils.CustomError
	UpdateEvent(*gin.Context, *models.Event) *utils.CustomError
	DeleteEvent(*gin.Context, *models.Event) error
}

func (eventService serviceImpl) GetEvents(context *gin.Context) ([]dto.EventDto, error) {
	events, err := eventService.Store.GetAllMongoEvents()
	eventsDto := utils.EventsDTOMapper(events)
	return eventsDto, err
}

func (eventService serviceImpl) GetEvent(context *gin.Context, ID string) (*models.Event, *utils.CustomError) {
	if ID == "" {
		log.Println("Invalid id, id cannot be empty")
		return nil, utils.NewCustomError(400, errors.New("invalid id, id cannot be empty"))
	}
	event, err := eventService.Store.GetMongoEventByID(ID)
	return event, utils.NewCustomError(404, err)
}

func (eventService serviceImpl) CreateEvents(context *gin.Context, events []models.Event) *utils.CustomError {
	for _, event := range events {
		check := utils.CheckValidLocation(event.Location)
		if !check {
			return utils.NewCustomError(400, errors.New("invalid location, please enter a proper location"))
		}
	}
	err := eventService.Store.SaveMongoEvent(events)
	return utils.NewCustomError(500, err)
}

func (eventService serviceImpl) UpdateEvent(context *gin.Context, event *models.Event) *utils.CustomError {
	check := utils.CheckValidLocation(event.Location)
	if !check {
		return utils.NewCustomError(400, errors.New("invalid location, please enter a proper location"))
	}
	err := eventService.Store.UpdateMongo(event)
	return utils.NewCustomError(500, err)
}

func (eventService serviceImpl) DeleteEvent(context *gin.Context, event *models.Event) error {
	err := eventService.Store.DeleteMongoEvent(event)
	return err
}
