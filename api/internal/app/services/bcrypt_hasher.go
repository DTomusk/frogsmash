package services

import "golang.org/x/crypto/bcrypt"

type BCryptHasher struct{}

func NewBCryptHasher() *BCryptHasher {
	return &BCryptHasher{}
}

func (h *BCryptHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (h *BCryptHasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
