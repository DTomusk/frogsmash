package factories

import (
	"frogsmash/internal/app/models"

	"github.com/google/uuid"
)

type Hasher interface {
	HashPassword(password string) (string, error)
}

type UserFactory struct {
	Hasher Hasher
}

func NewUserFactory(hasher Hasher) *UserFactory {
	return &UserFactory{
		Hasher: hasher,
	}
}

func (f *UserFactory) CreateNewUser(email, password string) (*models.User, error) {
	hashedPassword, err := f.Hasher.HashPassword(password)
	if err != nil {
		return nil, err
	}
	userID := uuid.New().String()
	return &models.User{
		ID:           userID,
		Email:        email,
		PasswordHash: hashedPassword,
	}, nil
}
