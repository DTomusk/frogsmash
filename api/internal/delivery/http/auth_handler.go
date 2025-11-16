package http

import (
	"database/sql"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/dto"

	"context"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	RegisterUser(email, password string, ctx context.Context, db repos.DBTX) error
	Login(email, password string, ctx context.Context, db repos.DBWithTxStarter) (string, *models.RefreshToken, *models.User, error)
	RefreshToken(refreshToken string, ctx context.Context, db repos.DBWithTxStarter) (string, *models.RefreshToken, *models.User, error)
}

type AuthHandler struct {
	AuthService AuthService
	db          *sql.DB
}

func NewAuthHandler(c *container.Container) *AuthHandler {
	return &AuthHandler{
		AuthService: c.AuthService,
		db:          c.DB,
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
	err := h.AuthService.RegisterUser(req.Email, req.Password, ctx.Request.Context(), h.db)
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
	jwt, refreshToken, user, err := h.AuthService.Login(req.Email, req.Password, ctx.Request.Context(), h.db)
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
	jwt, refreshToken, user, err := h.AuthService.RefreshToken(cookie, ctx.Request.Context(), h.db)
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
