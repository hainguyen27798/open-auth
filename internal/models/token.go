package models

import "time"

type Token struct {
	ID           string    `db:"id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	UserID       string    `db:"user_id"`
	Session      string    `db:"session"`
	RefreshToken string    `db:"refresh_token"`
}

type InsertNewTokenParams struct {
	UserID       string `db:"user_id"`
	Session      string `db:"session"`
	RefreshToken string `db:"refresh_token"`
}
