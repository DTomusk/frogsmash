package dto

const (
	AlreadyVerifiedCode = "ALREADY_VERIFIED"
	InvalidCodeCode     = "INVALID_VERIFICATION_CODE"
	VerifiedCode        = "USER_VERIFIED"
)

// VerificationRequest godoc
// @Description  Request payload for user verification
type VerificationRequest struct {
	Code string `json:"code" binding:"required"`
}

// ResendVerificationEmailRequest godoc
// @Description  Request payload for resending verification email
type ResendVerificationEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}
