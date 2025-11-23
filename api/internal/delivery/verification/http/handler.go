package http

import (
	"context"
	"errors"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/app/verification/services"
	"frogsmash/internal/container"
	sharedDto "frogsmash/internal/delivery/shared/dto"
	"frogsmash/internal/delivery/shared/utils"
	"frogsmash/internal/delivery/verification/dto"

	"github.com/gin-gonic/gin"
)

type VerificationService interface {
	ResendVerificationEmail(userID string, ctx context.Context, db shared.DBWithTxStarter) error
	VerifyUser(code, userID string, isVerified bool, ctx context.Context, db shared.DBWithTxStarter) error
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
		ctx.JSON(401, sharedDto.Response{
			Error: "Unauthorized",
			Code:  "UNAUTHORIZED",
		})
		return
	}

	err := h.verificationService.ResendVerificationEmail(user_id, ctx.Request.Context(), h.db)
	if err != nil {
		ctx.JSON(500, sharedDto.Response{
			Error: err.Error(),
			Code:  sharedDto.InternalServerErrorCode,
		})
		return
	}
	ctx.JSON(200, sharedDto.Response{
		Message: "Verification email resent successfully",
	})
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
		ctx.JSON(400, sharedDto.Response{
			Error: "Invalid request",
			Code:  sharedDto.InvalidRequestCode,
		})
		return
	}

	claims, hasClaims := utils.GetClaims(ctx)

	err := h.verificationService.VerifyUser(req.Code, claims.Sub, hasClaims && claims.IsVerified, ctx.Request.Context(), h.db)

	switch {
	case err == nil:
		ctx.JSON(200, sharedDto.Response{
			Message: "User verified successfully",
			Code:    dto.VerifiedCode,
		})
	case errors.Is(err, services.ErrInvalidVerificationCode):
		ctx.JSON(400, sharedDto.Response{
			Error: "Invalid verification code",
			Code:  dto.InvalidCodeCode,
		})
	case errors.Is(err, services.ErrAlreadyVerified):
		ctx.JSON(409, sharedDto.Response{
			Error: "User is already verified",
			Code:  dto.AlreadyVerifiedCode,
		})
	default:
		ctx.JSON(500, sharedDto.Response{
			Error: "Internal server error",
			Code:  sharedDto.InternalServerErrorCode,
		})
	}
}
