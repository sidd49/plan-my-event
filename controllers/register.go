package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRegisterController interface {
	RegisterForEvent(*gin.Context)
	CancelRegistration(*gin.Context)
}

// @Summary	Register for an event
// @Description	Registers an user to the Event
// @Tags	registration
// @Param id path string true "event id"
// @Success	201 {object} map[string]any
// @Failure	400 {object} map[string]any
// @Failure	404 {object} map[string]any
// @Failure	500 {object} map[string]any
// @Router /events/:id/register [post]
func (registerCont controllerImpl) RegisterForEvent(context *gin.Context) {
	// fetch userID of the logged in user
	userID := context.GetString("userID")
	eventID := context.Param("id")
	log.Println("Fetching the event with id : " + eventID)
	// Check if the event is present or not
	event, custErr := registerCont.service.GetEvent(context, eventID)
	if custErr.Err != nil {
		log.Println("Error in fetching the event : " + custErr.Err.Error())
		context.JSON(custErr.StatusCode, gin.H{"message": "Could not fetch the ID of event : " + custErr.Err.Error()})
		return
	}
	log.Println("Registering the user : " + userID)
	// Service layer call to register the user for the event
	err := registerCont.service.RegisterForEvent(context, event, userID)
	if err != nil {
		log.Println("Error in registration of the user : " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry ! Could not register you : " + err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered User for the Event !"})
}

// @Summary	Cancel Registration for an event
// @Description	Cancels Registration of an user to the Event
// @Tags	registration
// @Param id path string true "event id"
// @Success	202 {object} map[string]any
// @Failure	400 {object} map[string]any
// @Failure	500 {object} map[string]any
// @Router /events/:id/register [delete]
func (registerCont controllerImpl) CancelRegistration(context *gin.Context) {
	// Fetch userID from context of the logged in user
	userID := context.GetString("userID")
	eventID := context.Param("id")
	log.Println("Cancelling the registration of the event : " + eventID)
	// Service call to register the user for the event
	custErr := registerCont.service.CancelRegistration(context, eventID, userID)
	if custErr.Err != nil {
		log.Println("Error in cancelling the registration : " + custErr.Err.Error())
		context.JSON(custErr.StatusCode, gin.H{"message": custErr.Err.Error()})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "Event Registration cancelled successfully"})
}
