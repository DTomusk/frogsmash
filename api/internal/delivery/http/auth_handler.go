package http

import (
	"database/sql"
	"errors"
	"frogsmash/internal/app/auth/models"
	"frogsmash/internal/app/auth/services"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/dto"
	"frogsmash/internal/delivery/utils"

	"context"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Login(email, password string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *models.User, error)
	RefreshToken(refreshToken string, ctx context.Context, db shared.DBWithTxStarter) (string, *models.RefreshToken, *models.User, error)
}

type VerificationService interface {
	ResendVerificationEmail(userID string, ctx context.Context, db shared.DBWithTxStarter) error
	VerifyAnonymous(code string, ctx context.Context, db shared.DBWithTxStarter) error
	VerifyLoggedIn(code, userID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error
}

type UserService interface {
	RegisterUser(email, password string, ctx context.Context, db shared.DBWithTxStarter) error
}

type AuthHandler struct {
	authService         AuthService
	userService         UserService
	verificationService VerificationService
	db                  *sql.DB
}

func NewAuthHandler(c *container.Container) *AuthHandler {
	return &AuthHandler{
		authService:         c.Auth.AuthService,
		userService:         c.Auth.UserService,
		verificationService: c.Auth.VerificationService,
		db:                  c.InfraServices.DB,
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

// ResendVerificationEmail godoc
// @Summary      Resend verification email
// @Description  Resends the verification email to the user
// @Router       /resend-verification [post]
// @Accept       json
// @Produce      json
func (h *AuthHandler) ResendVerificationEmail(ctx *gin.Context) {
	user_id, ok := utils.GetUserID(ctx)
	if !ok {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	err := h.verificationService.ResendVerificationEmail(user_id, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Verification email resent successfully"})
}

// VerifyUser godoc
// @Summary      Verify user email
// @Description  Verifies the user's email using a verification code
// @Router       /verify [post]
// @Accept       json
// @Produce      json
// @Param        code  body  dto.VerificationRequest  true  "Verification code payload"
func (h *AuthHandler) VerifyUser(ctx *gin.Context) {
	var req dto.VerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil || req.Code == "" {
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	claims, hasClaims := utils.GetClaims(ctx)

	var err error

	if !hasClaims || claims.Sub == "" {
		// Verify anonymous user
		err = h.verificationService.VerifyAnonymous(req.Code, ctx.Request.Context(), h.db)
	} else {
		// Verify logged-in user
		err = h.verificationService.VerifyLoggedIn(req.Code, claims.Sub, claims.IsVerified, ctx.Request.Context(), h.db)
	}

	if err == nil {
		ctx.JSON(200, gin.H{
			"message": "User verified successfully",
			"code":    "USER_VERIFIED",
		})
		return
	}

	if errors.Is(err, services.ErrInvalidVerificationCode) {
		ctx.JSON(400, gin.H{
			"error": "Invalid verification code",
			"code":  "INVALID_VERIFICATION_CODE",
		})
		return
	}

	if errors.Is(err, services.ErrAlreadyVerified) {
		ctx.JSON(409, gin.H{
			"error": "User is already verified",
			"code":  "ALREADY_VERIFIED",
		})
		return
	}

	ctx.JSON(500, gin.H{"error": "Internal server error", "code": "INTERNAL_SERVER_ERROR"})
}
