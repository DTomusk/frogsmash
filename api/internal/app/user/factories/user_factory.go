package factories

import (
	"frogsmash/internal/app/user/models"

	"github.com/google/uuid"
)

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (f *UserFactory) CreateNewUser(email, hashedPassword string) (*models.User, error) {
	userID := uuid.New().String()
	return &models.User{
		ID:           userID,
		Email:        email,
		PasswordHash: hashedPassword,
	}, nil
}
