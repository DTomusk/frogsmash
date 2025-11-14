package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	secret               []byte
	tokenLifetimeMinutes int
}

func NewJwtService(secret []byte, tokenLifetimeMinutes int) *JwtService {
	return &JwtService{
		secret:               secret,
		tokenLifetimeMinutes: tokenLifetimeMinutes,
	}
}

func (s *JwtService) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Minute * time.Duration(s.tokenLifetimeMinutes)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *JwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return s.secret, nil
	})
}
