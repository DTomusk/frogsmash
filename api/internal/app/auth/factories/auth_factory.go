package factories

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"frogsmash/internal/app/auth/models"
	"time"
)

// TODO: consider where this should live
// IT's a small utility function that doesn't get used elsewhere and has no state, so it's fine here for now as a private function
// We probably don't need a service for it
func GenerateRefreshToken(userID string, tokenLifeTimeDays int) (*models.RefreshToken, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return &models.RefreshToken{
		Token:     fmt.Sprintf("%x", b),
		UserID:    userID,
		ExpiresAt: time.Now().UTC().Add(time.Duration(tokenLifeTimeDays) * 24 * time.Hour),
		MaxAge:    int64(tokenLifeTimeDays * 24 * 60 * 60),
		Revoked:   false,
	}, nil
}

func GenerateVerificationCode(userId string, codeLength, codeLifeTimeMinutes int) (*models.VerificationCode, error) {
	b := make([]byte, codeLength)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return &models.VerificationCode{
		UserID:    userId,
		Code:      base64.StdEncoding.EncodeToString(b),
		ExpiresAt: time.Now().UTC().Add(time.Duration(codeLifeTimeMinutes) * time.Minute),
	}, nil
}
