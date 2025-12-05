package services

import (
	"context"
	"fmt"

	"google.golang.org/api/idtoken"
)

type GoogleService interface {
	VerifyIDToken(idToken string, ctx context.Context) (string, error)
}

type googleService struct {
	// Add any dependencies needed for Google token verification
	ClientID string
}

func NewGoogleService(clientID string) GoogleService {
	return &googleService{
		ClientID: clientID,
	}
}

func (s *googleService) VerifyIDToken(idToken string, ctx context.Context) (string, error) {
	payload, err := idtoken.Validate(ctx, idToken, s.ClientID)
	if err != nil {
		return "", err
	}
	email, ok := payload.Claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("email not found in token")
	}
	return email, nil
}
