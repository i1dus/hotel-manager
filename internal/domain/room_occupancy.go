package domain

import (
	"time"
)

type RoomOccupancy struct {
	ID          int
	RoomNumber  string
	ClientID    int
	StartAt     time.Time
	EndAt       time.Time
	Description string
}
