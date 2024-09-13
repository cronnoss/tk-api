package models

import (
	"time"
)

type Show struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
