package models

import (
	"database/sql"
	"time"
)

type Role struct {
	ID          string         `db:"id"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	CanModify   int8           `db:"can_modify"`
}

type InsertNewRoleParams struct {
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type UpdateRoleParams struct {
	ID          *string `db:"id" json:"id"`
	Description *string `db:"description" json:"description,omitempty" attr:"description"`
}
