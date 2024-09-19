package models

import (
	"time"
)

type Event struct {
	ID          string
	Name        string `binding:"required"`
	Description string
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      string
}
