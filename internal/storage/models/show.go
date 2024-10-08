package models

import (
	"database/sql"
	"time"
)

type Show struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
