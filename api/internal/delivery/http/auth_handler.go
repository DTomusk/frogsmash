package http

import (
	"frogsmash/internal/app/auth/models"
	"frogsmash/internal/app/shared"
	user "frogsmash/internal/app/user/models"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/dto"
	"frogsmash/internal/delivery/utils"

	"context"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(email, password string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error)
	RefreshToken(refreshToken string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error)
}

type UserService interface {
	RegisterUser(email, password string, ctx context.Context, db shared.DBWithTxStarter) error
	GetUserByID(id string, ctx context.Context, db shared.DBWithTxStarter) (*user.User, error)
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

// GetMe godoc
// @Summary      Get current user
// @Description  Retrieves the currently authenticated user's information
// @Router       /me [get]
// @Produce      json
func (h *AuthHandler) GetMe(ctx *gin.Context) {
	sub, exists := utils.GetUserID(ctx)
	if !exists || sub == "" {
		ctx.JSON(401, dto.Response{
			Error: "Unauthorized",
			Code:  dto.UnauthorizedCode,
		})
		return
	}
	user, err := h.userService.GetUserByID(sub, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, dto.Response{
			Error: "Failed to retrieve user",
			Code:  dto.InternalServerErrorCode,
		})
		return
	}
	ctx.JSON(200, dto.UserResponse{
		ID:         user.ID,
		Email:      user.Email,
		IsVerified: user.IsVerified,
	})
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
		ctx.JSON(400, dto.Response{
			Error: "Invalid request payload",
			Code:  dto.InvalidRequestCode,
		})
		return
	}
	err := h.userService.RegisterUser(req.Email, req.Password, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, dto.Response{
			Error: err.Error(),
			Code:  dto.InternalServerErrorCode,
		})
		return
	}

	ctx.JSON(201, dto.Response{
		Message: "User registered successfully",
	})
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
		ctx.JSON(400, dto.Response{
			Error: "Invalid credentials",
			Code:  dto.InvalidRequestCode,
		})
		return
	}
	jwt, refreshToken, user, err := h.authService.Login(req.Email, req.Password, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, dto.Response{
			Error: "Invalid credentials",
			Code:  dto.InternalServerErrorCode,
		})
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
		ctx.JSON(400, dto.Response{
			Error: "Refresh token not provided",
			Code:  dto.InvalidRequestCode,
		})
		return
	}
	jwt, refreshToken, user, err := h.authService.RefreshToken(cookie, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, dto.Response{
			Error: err.Error(),
			Code:  dto.InternalServerErrorCode,
		})
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
