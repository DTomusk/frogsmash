package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateToken(userID string, isVerified bool) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
	secret               []byte
	tokenLifetimeMinutes int
}

func NewJwtService(secret []byte, tokenLifetimeMinutes int) TokenService {
	return &jwtService{
		secret:               secret,
		tokenLifetimeMinutes: tokenLifetimeMinutes,
	}
}

func (s *jwtService) GenerateToken(userID string, isVerified bool) (string, error) {
	claims := jwt.MapClaims{
		"sub":         userID,
		"exp":         time.Now().UTC().Add(time.Minute * time.Duration(s.tokenLifetimeMinutes)).Unix(),
		"is_verified": isVerified,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return s.secret, nil
	})
}
