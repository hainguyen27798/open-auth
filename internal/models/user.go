package models

import (
	"database/sql"
	"time"
)

type UsersScope string

const (
	UsersScopeUser  UsersScope = "user"
	UsersScopeAdmin UsersScope = "admin"
)

type UsersStatus string

const (
	UsersStatusActive   UsersStatus = "active"
	UsersStatusInActive UsersStatus = "inActive"
	UsersStatusRequest  UsersStatus = "request"
)

type UsersSocialProvider string

const (
	UsersSocialProviderBasic    UsersSocialProvider = "basic"
	UsersSocialProviderGoogle   UsersSocialProvider = "google"
	UsersSocialProviderLinkedin UsersSocialProvider = "linkedin"
)

type User struct {
	ID               string              `db:"id"`
	CreatedAt        time.Time           `db:"created_at"`
	UpdatedAt        time.Time           `db:"updated_at"`
	Name             string              `db:"name"`
	Email            string              `db:"email"`
	Password         sql.NullString      `db:"password"`
	Status           UsersStatus         `db:"status"`
	SocialProvider   UsersSocialProvider `db:"social_provider"`
	Image            sql.NullString      `db:"image"`
	Verify           int8                `db:"verify"`
	VerificationCode sql.NullString      `db:"verification_code"`
	RoleID           sql.NullString      `db:"role_id"`
	Scope            UsersScope          `db:"scope"`
}

type InsertNewUserParams struct {
	Name             string `db:"name"`
	Email            string `db:"email"`
	Password         string `db:"password"`
	VerificationCode string `db:"verification_code"`
}

type InsertSuperUserParams struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}
