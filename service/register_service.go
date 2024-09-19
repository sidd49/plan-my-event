package service

import (
	"errors"

	models "example.com/plan-my-event/models"
	"example.com/plan-my-event/utils"
	"github.com/gin-gonic/gin"
)

type IRegisterService interface {
	RegisterForEvent(*gin.Context, *models.Event, string) error
	CancelRegistration(*gin.Context, string, string) *utils.CustomError
}

func (registerService serviceImpl) RegisterForEvent(ctx *gin.Context, event *models.Event, userID string) error {
	err := registerService.Store.MongoRegister(event, userID)
	return err
}

func (registerService serviceImpl) CancelRegistration(ctx *gin.Context, eventID string, userID string) *utils.CustomError {
	count, err := registerService.Store.CancelMongoRegistration(eventID, userID)
	if (count == 0) && (err == nil) {
		return utils.NewCustomError(400, errors.New("could not get the registration of the event with the user from db"))
	}
	return utils.NewCustomError(500, err)
}
