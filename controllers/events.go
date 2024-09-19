package controllers

import (
	"log"
	"net/http"

	models "example.com/plan-my-event/models"
	"github.com/gin-gonic/gin"
)

type IEventsController interface {
	GetEvents(*gin.Context)
	GetEvent(*gin.Context)
	CreateEvent(*gin.Context)
	UpdateEvent(*gin.Context)
	DeleteEvent(*gin.Context)
}

// @Summary	Get all events
// @Description	Displays all the events present in the app
// @Tags	events
// @Success	200 {array} dto.EventDto
// @Failure	500 {object} map[string]any
// @Router /events [get]
func (eventCont controllerImpl) GetEvents(context *gin.Context) {
	log.Println("Retreiving all the events")
	//fetch events from service layer
	events, err := eventCont.service.GetEvents(context)
	if err != nil {
		log.Println("Error in getting all events: " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if events == nil {
		context.JSON(http.StatusOK, gin.H{"message": "No events created yet in the application"})
		return
	}
	context.JSON(http.StatusOK, events)
}

// @Summary Get particular event by id
// @Description Get the one event
// @Tags events
// @Param id path string true "event id"
// @Success 200 {object} models.Event
// @Failure 400 {object} map[string]any
// @Failure 404 {object} map[string]any
// @Router /events/:id [get]
func (eventCont controllerImpl) GetEvent(context *gin.Context) {
	eventId := context.Param("id")
	log.Println("Fetching event with id : " + eventId)
	//fetch event with parameter id from service layer
	event, custErr := eventCont.service.GetEvent(context, eventId)
	if custErr.Err != nil {
		log.Println("Error in finding the event with id : " + eventId + " Error : " + custErr.Err.Error())
		context.JSON(custErr.StatusCode, gin.H{"message": "Could not find the event with id : " + eventId})
		return
	}
	context.JSON(http.StatusOK, event)
}

// @Summary Create new event
// @Description Create a new event
// @Tags events
// @Param event body models.Event true "Event Object"
// @Success 201 {object} models.Event
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /events [post]
func (eventCont controllerImpl) CreateEvent(context *gin.Context) {
	var event models.Event
	// bind the payload with the event model
	err := context.ShouldBindJSON(&event)
	if err != nil {
		log.Println("Error in parsing request for event : " + err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message": "Events could not be parsed"})
		return
	}
	//take the user id of the logged in user to save for the event
	event.UserID = context.GetString("userID")
	log.Println("Started creating event with id : " + event.ID)
	var events []models.Event
	events = append(events, event)
	// Service layer call to save the event
	custErr := eventCont.service.CreateEvents(context, events)
	if custErr.Err != nil {
		log.Println("Error in creating the event : " + custErr.Err.Error())
		context.JSON(custErr.StatusCode, gin.H{"message": custErr.Err.Error()})
		return
	}
	context.JSON(http.StatusCreated, event)
}

// @Summary Update an event
// @Description Update event
// @Tags events
// @Param event body models.Event true "Event Object"
// @Success 200 {object} map[string]any
// @Failure 400 {object} map[string]any
// @Failure 401 {object} map[string]any
// @Failure 404 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /events/:id [put]
func (eventCont controllerImpl) UpdateEvent(context *gin.Context) {
	eventId := context.Param("id")
	log.Println("Checking if the event is present or not")
	// Check first if the event is present or not
	event, custErr := eventCont.service.GetEvent(context, eventId)
	if custErr.Err != nil {
		log.Println("Error in finding the event from db : " + custErr.Err.Error())
		context.JSON(custErr.StatusCode, gin.H{"message": "Could not fetch the event", "error": custErr.Err.Error()})
		return
	}
	// check if the user who is updating the event is the creator or not
	if event.UserID != context.GetString("userID") {
		log.Println("User not authorized to update the event")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "This user not authorized to update this event"})
		return
	}
	var updatedEvent models.Event
	// bind the payload with event model
	err := context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		log.Println("Error in parsing the request : " + err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedEvent.ID = event.ID
	updatedEvent.UserID = event.UserID
	log.Println("Started updating the event with id : " + updatedEvent.ID)
	// Service layer call to update the event
	custErr = eventCont.service.UpdateEvent(context, &updatedEvent)
	if custErr.Err != nil {
		log.Println("Error in updating the event : " + custErr.Err.Error())
		context.JSON(custErr.StatusCode, gin.H{"message": custErr.Err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

// @Summary Delete event
// @Description Delete an event
// @Tags events
// @Param id path string true "Event id"
// @Success 204
// @Failure 400 {object} map[string]any
// @Failure 401 {object} map[string]any
// @Failure 404 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /events/:id [delete]
func (eventCont controllerImpl) DeleteEvent(context *gin.Context) {
	eventId := context.Param("id")
	log.Println("Fetching the event to delete")
	// Check if the event exists or not
	event, custErr := eventCont.service.GetEvent(context, eventId)
	if custErr.Err != nil {
		log.Println("Error in fetching the event with the given id : " + custErr.Err.Error())
		context.JSON(custErr.StatusCode, gin.H{"message": "Could not fetch the event"})
		return
	}
	// Check if the user who created the event is deleting or not
	if event.UserID != context.GetString("userID") {
		log.Println("User not authorized to delete the event")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized to delete the event"})
		return
	}
	log.Println("Deleting the event with id : " + eventId)
	// Service layer call to delete events
	err := eventCont.service.DeleteEvent(context, event)
	if err != nil {
		log.Println("Error in deleting the event : " + err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{})
}
