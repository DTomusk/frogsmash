package models

import "time"

type VerificationCode struct {
	Code      string
	UserID    string
	ExpiresAt time.Time
}
