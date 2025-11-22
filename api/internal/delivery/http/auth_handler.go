package http

import (
	"frogsmash/internal/app/auth/models"
	"frogsmash/internal/app/shared"
	user "frogsmash/internal/app/user/models"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/dto"

	"context"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(email, password string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error)
	RefreshToken(refreshToken string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error)
}

type UserService interface {
	RegisterUser(email, password string, ctx context.Context, db shared.DBWithTxStarter) error
}

type AuthHandler struct {
	authService AuthService
	userService UserService
	db          shared.DBWithTxStarter
}

func NewAuthHandler(c *container.Container) *AuthHandler {
	return &AuthHandler{
		authService: c.Auth.AuthService,
		db:          c.InfraServices.DB,
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Registers a new user with email and password
// @Router       /register [post]
// @Accept       json
// @Produce      json
// @Param        user  body  dto.UserRegistrationRequest  true  "User registration payload"
func (h *AuthHandler) Register(ctx *gin.Context) {
	var req dto.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.userService.RegisterUser(req.Email, req.Password, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"message": "User registered successfully"})
}

// Login godoc
// @Summary      User login
// @Description  Logs in a user with email and password
// @Router       /login [post]
// @Accept       json
// @Produce      json
// @Param        user  body  dto.UserLoginRequest  true  "User login payload"
// @Success      200   {object}  dto.UserLoginResponse
func (h *AuthHandler) Login(ctx *gin.Context) {
	// Implementation for user login
	var req dto.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid credentials"})
		return
	}
	jwt, refreshToken, user, err := h.authService.Login(req.Email, req.Password, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Invalid credentials"})
		return
	}

	setRefreshTokenCookie(ctx, refreshToken)

	res := dto.UserLoginResponse{
		JWT: jwt,
		User: dto.UserResponse{
			ID:         user.ID,
			Email:      user.Email,
			IsVerified: user.IsVerified,
		},
	}
	ctx.JSON(200, res)
}

// RefreshToken godoc
// @Summary      Refresh JWT token
// @Description  Refreshes the JWT token using a refresh token
// @Router       /refresh-token [post]
// @Accept       json
// @Produce      json
// @Success      200    {object}  dto.UserLoginResponse
func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Refresh token cookie missing"})
		return
	}
	jwt, refreshToken, user, err := h.authService.RefreshToken(cookie, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	setRefreshTokenCookie(ctx, refreshToken)

	res := dto.UserLoginResponse{
		JWT: jwt,
		User: dto.UserResponse{
			ID:         user.ID,
			Email:      user.Email,
			IsVerified: user.IsVerified,
		},
	}
	ctx.JSON(200, res)
}

func setRefreshTokenCookie(ctx *gin.Context, refreshToken *models.RefreshToken) {
	ctx.SetCookie(
		"refresh_token",
		refreshToken.Token,
		int(refreshToken.MaxAge),
		"/",
		"",
		// TODO: move to https in production
		false,
		true,
	)
}
