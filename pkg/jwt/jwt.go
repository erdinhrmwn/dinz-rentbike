package jwt

import (
	"dinz-rentbike/internal/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	UserID uint `json:"user_id"`
}

type AuthManager struct {
	secret string
}

func NewAuthManager(cfg *config.JwtConfig) *AuthManager {
	return &AuthManager{
		secret: cfg.Secret,
	}
}

func (am *AuthManager) GenerateToken(userID uint) (*string, error) {
	claims := JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(am.secret))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (am *AuthManager) VerifyToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(am.secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
