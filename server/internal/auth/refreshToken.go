package auth

import (
	"crypto/rand"
	"time"

	"github.com/sebasvelasco353/nummus/server/internal/config"
)

// creates a new random string with high entropy
func GenerateRefreshToken() string {
	return rand.Text()
}

func PersistRefreshToken(token string, userId string, expDate time.Time) (string, error) {
	var err error
	var newTokenId string

	query := "INSERT INTO refresh_tokens (refresh_token_hash, user_id, exp_date, is_used) VALUES ($1, $2, $3, $4) RETURNING refresh_token_id"
	err = config.DB.QueryRow(query, token, userId, expDate, false).Scan(&newTokenId)
	if err != nil {
		return "", err
	}
	return newTokenId, err
}
