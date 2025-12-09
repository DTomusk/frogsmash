package http

import (
	"frogsmash/internal/app/auth/models"
	"frogsmash/internal/app/shared"
	user "frogsmash/internal/app/user/models"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/auth/dto"
	sharedDto "frogsmash/internal/delivery/shared/dto"
	"frogsmash/internal/delivery/shared/utils"

	"context"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(email, password, tenantID string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error)
	Logout(refreshToken string, ctx context.Context, db shared.DBWithTxStarter) error
	RefreshToken(refreshToken, tenantID string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error)
	Register(email, password, tenantID string, ctx context.Context, db shared.DBWithTxStarter) error
	GoogleLogin(idToken, tenantID string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *user.User, error)
}

type UserService interface {
	GetUserByUserID(userID, tenantID string, ctx context.Context, db shared.DBTX) (*user.User, error)
}

type AuthHandler struct {
	authService AuthService
	userService UserService
	db          shared.DBWithTxStarter
}

func NewAuthHandler(c *container.APIContainer) *AuthHandler {
	return &AuthHandler{
		authService: c.Auth.AuthService,
		userService: c.User.UserService,
		db:          c.InfraServices.DB,
	}
}

// GetMe godoc
// @Summary      Get current user
// @Description  Retrieves the currently authenticated user's information
// @Router       /auth/me [get]
// @Produce      json
func (h *AuthHandler) GetMe(ctx *gin.Context) {
	tenantID, exists := ctx.Get("tenant_id")
	if !exists {
		ctx.JSON(500, sharedDto.Response{
			Error: "Tenant ID not found in context",
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}
	sub, exists := utils.GetUserID(ctx)
	if !exists || sub == "" {
		ctx.JSON(401, sharedDto.Response{
			Error: "Unauthorized",
			Code:  sharedDto.UnauthorizedCode,
		})
		return
	}
	user, err := h.userService.GetUserByUserID(sub, tenantID.(string), ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: "Failed to retrieve user",
			Code:  sharedDto.InternalServerErrorCode,
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
// @Router       /auth/register [post]
// @Accept       json
// @Produce      json
// @Param        user  body  dto.UserRegistrationRequest  true  "User registration payload"
func (h *AuthHandler) Register(ctx *gin.Context) {
	tenantID, exists := ctx.Get("tenant_id")
	if !exists {
		ctx.JSON(500, sharedDto.Response{
			Error: "Tenant ID not found in context",
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}
	var req dto.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, sharedDto.Response{
			Error: "Invalid request payload",
			Code:  sharedDto.InvalidRequestCode,
		})
		return
	}
	err := h.authService.Register(req.Email, req.Password, tenantID.(string), ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: err.Error(),
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}

	ctx.JSON(201, sharedDto.Response{
		Message: "User registered successfully",
	})
}

// Login godoc
// @Summary      User login
// @Description  Logs in a user with email and password
// @Router       /auth/login [post]
// @Accept       json
// @Produce      json
// @Param        user  body  dto.UserLoginRequest  true  "User login payload"
// @Success      200   {object}  dto.UserLoginResponse
func (h *AuthHandler) Login(ctx *gin.Context) {
	tenantID, exists := ctx.Get("tenant_id")
	if !exists {
		ctx.JSON(500, sharedDto.Response{
			Error: "Tenant ID not found in context",
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}
	var req dto.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, sharedDto.Response{
			Error: "Invalid credentials",
			Code:  sharedDto.InvalidRequestCode,
		})
		return
	}
	jwt, refreshToken, user, err := h.authService.Login(req.Email, req.Password, tenantID.(string), ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: "Invalid credentials",
			Code:  sharedDto.InternalServerErrorCode,
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

// Logout godoc
// @Summary      User logout
// @Description  Logs out a user by clearing the refresh token cookie
// @Router       /auth/logout [post]
// @Accept       json
// @Produce      json
func (h *AuthHandler) Logout(ctx *gin.Context) {
	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(400, sharedDto.Response{
			Error: "Refresh token not provided",
			Code:  sharedDto.InvalidRequestCode,
		})
		return
	}
	// Ignore error to ensure logout proceeds
	// TODO: log error if needed
	_ = h.authService.Logout(cookie, ctx.Request.Context(), h.db)

	ctx.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		// TODO: move to https in production
		false,
		true,
	)
	ctx.JSON(200, sharedDto.Response{
		Message: "User logged out successfully",
	})
}

// RefreshToken godoc
// @Summary      Refresh JWT token
// @Description  Refreshes the JWT token using a refresh token
// @Router       /auth/refresh-token [post]
// @Accept       json
// @Produce      json
// @Success      200    {object}  dto.UserLoginResponse
func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	tenantID, exists := ctx.Get("tenant_id")
	if !exists {
		ctx.JSON(500, sharedDto.Response{
			Error: "Tenant ID not found in context",
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}
	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(400, sharedDto.Response{
			Error: "Refresh token not provided",
			Code:  sharedDto.InvalidRequestCode,
		})
		return
	}
	jwt, refreshToken, user, err := h.authService.RefreshToken(cookie, tenantID.(string), ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: err.Error(),
			Code:  sharedDto.InternalServerErrorCode,
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
		true,
		true,
	)
}

func (h *AuthHandler) GoogleLogin(ctx *gin.Context) {
	tenantID, exists := ctx.Get("tenant_id")
	if !exists {
		ctx.JSON(500, sharedDto.Response{
			Error: "Tenant ID not found in context",
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}
	var req dto.GoogleLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, sharedDto.Response{
			Error: "Invalid request payload",
			Code:  sharedDto.InvalidRequestCode,
		})
		return
	}
	jwt, refreshToken, user, err := h.authService.GoogleLogin(req.IdToken, tenantID.(string), ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: err.Error(),
			Code:  sharedDto.InternalServerErrorCode,
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
