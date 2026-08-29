package auth

import (
	"strconv"
	"time"

	"github.com/sebasvelasco353/nummus/server/internal/config"
)

func PersistRefreshToken(token string, userId string) (string, error) {
	var err error
	var newTokenId string
	expDays, err := strconv.ParseInt(config.ServerCfg.RefreshExpiryDays, 10, 64)
	if err != nil {
		return "", err
	}
	expDate := time.Now().UTC().AddDate(0, 0, int(expDays))

	query := "INSERT INTO refresh_tokens (refresh_token_hash, user_id, exp_date, is_used) VALUES ($1, $2, $3, $4) RETURNING refresh_token_id"
	err = config.DB.QueryRow(query, token, userId, expDate, false).Scan(&newTokenId)
	if err != nil {
		return "", err
	}
	return newTokenId, err
}
