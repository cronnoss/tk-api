package models

import (
	"time"
)

type Place struct {
	ID          int64
	X           float64
	Y           float64
	Width       float64
	Height      float64
	IsAvailable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
