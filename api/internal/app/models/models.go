package models

import "time"

type Item struct {
	ID       string
	Name     string
	ImageURL string
	Score    float64
}

// TODO: JSON isn't a model concern, consider moving to dto package
type LeaderboardItem struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Score     float64 `json:"score"`
	ImageURL  string  `json:"image_url"`
	Rank      int     `json:"rank"`
	CreatedAt string  `json:"created_at"`
	License   string  `json:"license"`
}

type Event struct {
	ID       string
	WinnerID string
	LoserID  string
}

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
