package models

import (
	"database/sql"
	"time"
)

type Event struct {
	ID        int64        `db:"id"`
	ShowID    int64        `db:"show_id"`
	Date      string       `db:"date"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
