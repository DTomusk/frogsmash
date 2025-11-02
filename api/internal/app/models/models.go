package models

type Item struct {
	ID       string
	Name     string
	ImageURL string
	Score    float64
}

type Event struct {
	ID       string
	WinnerID string
	LoserID  string
}
