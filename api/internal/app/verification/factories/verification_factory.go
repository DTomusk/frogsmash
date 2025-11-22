package factories

import (
	"crypto/rand"
	"encoding/base64"
	"frogsmash/internal/app/verification/models"
	"time"
)

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
