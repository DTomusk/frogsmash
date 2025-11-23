package dto

import "math"

const (
	InvalidRequestCode      = "INVALID_REQUEST"
	UnauthorizedCode        = "UNAUTHORIZED"
	InternalServerErrorCode = "INTERNAL_SERVER_ERROR"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    string `json:"code,omitempty"`
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
