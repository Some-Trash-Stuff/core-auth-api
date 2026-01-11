package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	issuer     string
	secretKey  []byte
	expMinutes int
}

func NewAuthService(issuer string, secret string, expMinutes int) *AuthService {
	return &AuthService{
		issuer:     issuer,
		secretKey:  []byte(secret),
		expMinutes: expMinutes,
	}
}

func (s *AuthService) GeneratePublicToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"iss": s.issuer,
		"sub": userID,
		"aud": "public",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Minute * time.Duration(s.expMinutes)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}
