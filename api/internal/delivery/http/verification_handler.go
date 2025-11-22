package http

import (
	"context"
	"errors"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/app/verification/services"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/dto"
	"frogsmash/internal/delivery/utils"

	"github.com/gin-gonic/gin"
)

type VerificationService interface {
	ResendVerificationEmail(userID string, ctx context.Context, db shared.DBWithTxStarter) error
	VerifyAnonymous(code string, ctx context.Context, db shared.DBWithTxStarter) error
	VerifyLoggedIn(code, userID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error
}

type VerificationHandler struct {
	verificationService VerificationService
	db                  shared.DBWithTxStarter
}

func NewVerificationHandler(c *container.Container) *VerificationHandler {
	return &VerificationHandler{
		verificationService: c.Verification.VerificationService,
		db:                  c.InfraServices.DB,
	}
}

// ResendVerificationEmail godoc
// @Summary      Resend verification email
// @Description  Resends the verification email to the user
// @Router       /resend-verification [post]
// @Accept       json
// @Produce      json
func (h *VerificationHandler) ResendVerificationEmail(ctx *gin.Context) {
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
func (h *VerificationHandler) VerifyUser(ctx *gin.Context) {
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

	switch {
	case err == nil:
		ctx.JSON(200, gin.H{
			"message": "User verified successfully",
			"code":    "USER_VERIFIED",
		})
	case errors.Is(err, services.ErrInvalidVerificationCode):
		ctx.JSON(400, gin.H{
			"error": "Invalid verification code",
			"code":  "INVALID_VERIFICATION_CODE",
		})
	case errors.Is(err, services.ErrAlreadyVerified):
		ctx.JSON(409, gin.H{
			"error": "User is already verified",
			"code":  "ALREADY_VERIFIED",
		})
	default:
		ctx.JSON(500, gin.H{
			"error": "Internal server error",
			"code":  "INTERNAL_SERVER_ERROR",
		})
	}
}
