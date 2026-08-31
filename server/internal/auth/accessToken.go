package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/sebasvelasco353/nummus/server/internal/config"
)

type AccessTokenClaims struct {
	Email  string           `json:"email"`
	UserID string           `json:"userId"`
	Exp    *jwt.NumericDate `json:"exp"`
	jwt.RegisteredClaims
}

// generates a new signed JWT token
func GenerateAccessToken(claims AccessTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.ServerCfg.JWTSecret))
}

func ValidateAccessToken(token string) (string, error) {
	return "", nil
}
