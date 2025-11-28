package models

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

type ImageUpload struct {
	ID         string
	UserID     string
	URL        string
	UploadedAt string
}
