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

type InsertNewPermissionParams struct {
	ServiceName string `db:"serviceName" json:"serviceName"`
	Resource    string `db:"resource" json:"resource"`
	Action      string `db:"action" json:"action"`
	Attributes  string `db:"attributes" json:"attributes"`
	Description string `db:"description" json:"description"`
}

type UpdatePermissionParams struct {
	ID          *string `db:"id" json:"id"`
	ServiceName *string `db:"serviceName" json:"serviceName,omitempty" attr:"service_name"`
	Resource    *string `db:"resource" json:"resource,omitempty" attr:"resource"`
	Action      *string `db:"action" json:"action,omitempty" attr:"action"`
	Attributes  *string `db:"attributes" json:"attributes,omitempty" attr:"attributes"`
	Description *string `db:"description" json:"description,omitempty" attr:"description"`
}
