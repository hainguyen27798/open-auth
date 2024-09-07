// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type UsersSocialProvider string

const (
	UsersSocialProviderGoogle   UsersSocialProvider = "google"
	UsersSocialProviderLinkedin UsersSocialProvider = "linkedin"
)

func (e *UsersSocialProvider) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersSocialProvider(s)
	case string:
		*e = UsersSocialProvider(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersSocialProvider: %T", src)
	}
	return nil
}

type NullUsersSocialProvider struct {
	UsersSocialProvider UsersSocialProvider
	Valid               bool // Valid is true if UsersSocialProvider is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersSocialProvider) Scan(value interface{}) error {
	if value == nil {
		ns.UsersSocialProvider, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersSocialProvider.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersSocialProvider) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersSocialProvider), nil
}

type UsersStatus string

const (
	UsersStatusActive   UsersStatus = "active"
	UsersStatusInActive UsersStatus = "inActive"
	UsersStatusRequest  UsersStatus = "request"
)

func (e *UsersStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersStatus(s)
	case string:
		*e = UsersStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersStatus: %T", src)
	}
	return nil
}

type NullUsersStatus struct {
	UsersStatus UsersStatus
	Valid       bool // Valid is true if UsersStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UsersStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersStatus), nil
}

type Permission struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ServiceName string
	Resource    string
	Action      string
	Attributes  string
	Description sql.NullString
}

type RefreshTokensUsed struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	TokenID      string
	RefreshToken string
}

type Token struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       string
	Session      string
	RefreshToken string
}

type User struct {
	ID               string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Name             string
	Email            string
	Password         sql.NullString
	Status           UsersStatus
	SocialProvider   NullUsersSocialProvider
	Image            sql.NullString
	Verify           int8
	VerificationCode sql.NullString
}
