package model

import (
	"time"
)

// Event - структура события
type Event struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	EventDate   time.Time `json:"event_date"`
	Description string    `json:"description"`
}
