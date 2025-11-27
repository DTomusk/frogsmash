package dto

type GetComparisonItemsResponse struct {
	LeftItem  ItemDTO `json:"left_item"`
	RightItem ItemDTO `json:"right_item"`
}

type ItemDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

// CompareRequest godoc
// @Description  Request payload for comparing two items
type CompareRequest struct {
	WinnerId string `json:"winner_id"`
	LoserId  string `json:"loser_id"`
}

type GetLatestSubmissionResponse struct {
	UploadedAt string `json:"uploaded_at"`
}
