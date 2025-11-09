package models

type Item struct {
	ID       string
	Name     string
	ImageURL string
	Score    float64
}

type LeaderboardItem struct {
	ID       string
	Name     string
	Score    float64
	ImageURL string
	Rank     int
}

type Event struct {
	ID       string
	WinnerID string
	LoserID  string
}
