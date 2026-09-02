package auth

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"time"

	"github.com/sebasvelasco353/nummus/server/internal/config"
)

type RefreshTokenType struct {
	RefreshTokenID   string
	RefreshTokenHash string
	UserID           string
	IsUsed           bool
	ExpDate          time.Time
}

// creates a new random string with high entropy
func GenerateRefreshToken() string {
	return rand.Text()
}

// stores a refresh token on the database
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

// searches a specific refresh token's data on the database
func FetchRefreshToken(token string) (RefreshTokenType, error) {
	var err error
	var row RefreshTokenType

	query := "SELECT refresh_token_id, refresh_token_hash, user_id, exp_date, is_used FROM refresh_tokens WHERE refresh_token_hash = $1;"
	err = config.DB.QueryRow(query, token).Scan(&row.RefreshTokenID, &row.RefreshTokenHash, &row.UserID, &row.ExpDate, &row.IsUsed)

	if err == sql.ErrNoRows {
		return RefreshTokenType{}, fmt.Errorf("no token found")
	} else if err != nil {
		return RefreshTokenType{}, fmt.Errorf("there was an error while fetching the token from the database")
	}
	return row, nil
}
