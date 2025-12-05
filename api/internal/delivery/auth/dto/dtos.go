package dto

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
	IsVerified bool   `json:"isVerified"`
}

// GoogleLoginRequest godoc
// @Description  Request payload for Google login
type GoogleLoginRequest struct {
	IdToken string `json:"idToken" binding:"required"`
}
