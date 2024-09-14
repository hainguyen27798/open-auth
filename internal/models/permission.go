package models

import (
	"database/sql"
	"time"
)

type Permission struct {
	ID          string         `db:"id"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
	ServiceName string         `db:"service_name"`
	Resource    string         `db:"resource"`
	Action      string         `db:"action"`
	Attributes  string         `db:"attributes"`
	Description sql.NullString `db:"description"`
}
