package dto

import "math"

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

type PagedResponse[T any] struct {
	Data       []T `json:"data"`
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}

func NewPagedResponse[T any](data []T, total, page, limit int) PagedResponse[T] {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return PagedResponse[T]{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}

// UserRegistrationRequest godoc
// @Description  Request payload for user registration
type UserRegistrationRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=12"`
}

// UserLoginRequest godoc
// @Description  Request payload for user login
type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserLoginResponse godoc
// @Description  Response payload for user login
type UserLoginResponse struct {
	JWT  string       `json:"jwt"`
	User UserResponse `json:"user"`
}

// UserResponse godoc
// @Description  User data returned with login response
type UserResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}
