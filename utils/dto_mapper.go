package utils

import (
	"log"

	"example.com/plan-my-event/dto"
	models "example.com/plan-my-event/models"
)

func EventDTOMapper(event models.Event) (eventDto dto.EventDto) {
	log.Println("Mapping started for event with id : " + event.ID)
	eventDto.Name = event.Name
	eventDto.Description = event.Description
	eventDto.Location = event.Location
	eventDto.ID = event.ID
	eventDto.DateTime = event.DateTime
	return eventDto
}

func EventsDTOMapper(events []models.Event) (eventsDto []dto.EventDto) {
	for _, event := range events {
		eventsDto = append(eventsDto, EventDTOMapper(event))
	}
	return eventsDto
}
