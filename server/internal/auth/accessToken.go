package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sebasvelasco353/nummus/server/internal/config"
)

type AccessTokenClaims struct {
	Email  string `json:"email"`
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

// generates a new signed JWT token
func GenerateAccessToken(claims AccessTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.ServerCfg.JWTSecret))
}

// Validates the token is valid
func ValidateAccessToken(token string) (AccessTokenClaims, error) {
	var claims AccessTokenClaims
	var err error
	decodedToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			// if method of token cant be casted to a SigningMethodHMAC
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// if it is the right method, return the secret
		return []byte(config.ServerCfg.JWTSecret), nil
	})
	if err != nil {
		return AccessTokenClaims{}, err
	}
	if !decodedToken.Valid {
		return AccessTokenClaims{}, fmt.Errorf("access token is not valid")
	}
	return claims, err
}
