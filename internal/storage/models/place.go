package models

import (
	"database/sql"
	"time"
)

type Place struct {
	ID          int64        `db:"id"`
	X           float64      `db:"x"`
	Y           float64      `db:"y"`
	Width       float64      `db:"width"`
	Height      float64      `db:"height"`
	IsAvailable bool         `db:"is_available"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
}
