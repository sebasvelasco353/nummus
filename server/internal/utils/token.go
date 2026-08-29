package utils

import (
	"crypto/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sebasvelasco353/nummus/server/internal/config"
)

func GenerateAccessToken(email string, userId string) (string, error) {
	expHours, err := strconv.ParseInt(config.ServerCfg.JWTExpiryHours, 10, 64)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * time.Duration(expHours)),
	})
	return token.SignedString([]byte(config.ServerCfg.JWTSecret))
}

func GenerateRefreshToken() string {
	return rand.Text()
}

// TODO: Add token validation utility
func ValidateToken(token string) (bool, error) {
	return false, nil
}
