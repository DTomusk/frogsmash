package models

import "time"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    string
	IsVerified   bool
}

// TODO: it might be silly having both ExpiresAt and MaxAge, but for now it's convenient
type RefreshToken struct {
	Token     string
	UserID    string
	MaxAge    int64
	ExpiresAt time.Time
	CreatedAt time.Time
	Revoked   bool
}

type VerificationCode struct {
	Code      string
	UserID    string
	ExpiresAt time.Time
}

// TODO: consider moving
type Claims struct {
	Sub        string `json:"sub"`
	IsVerified bool   `json:"is_verified"`
}
