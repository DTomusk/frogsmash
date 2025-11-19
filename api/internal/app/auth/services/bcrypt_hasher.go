package services

import "golang.org/x/crypto/bcrypt"

type bCryptHasher struct{}

func NewBCryptHasher() Hasher {
	return bCryptHasher{}
}

func (h bCryptHasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (h bCryptHasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
