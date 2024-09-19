package dto

import "time"

type EventDto struct {
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	ID          string
}
