package domain

import (
	"time"
)

type RoomOccupancy struct {
	ID          int
	RoomNumber  string
	Passport    string
	StartAt     time.Time
	EndAt       *time.Time
	Description string
}
