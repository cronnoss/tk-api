package models

import (
	"time"
)

type Event struct {
	ID        int64
	ShowID    int64
	Date      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
