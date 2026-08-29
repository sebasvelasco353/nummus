package users

import (
	"strconv"
	"time"

	"github.com/sebasvelasco353/nummus/server/internal/auth"
	"github.com/sebasvelasco353/nummus/server/internal/config"
	"github.com/sebasvelasco353/nummus/server/internal/utils"
)

type User struct {
	ID       string
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Name     string `binding:"required"`
	LastName string
}

type UserLogin struct {
	ID       string
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type loginResult struct {
	UserId              string
	AccessToken         string
	RefreshToken        string
	RefreshTokenExpDate time.Time
}

func (u User) SignUp() (string, error) {
	var err error
	var newUserId string

	query := "INSERT INTO users (email, password, name, last_name) VALUES ($1, $2, $3, $4) RETURNING user_id"
	u.Password, err = auth.HashPassword(u.Password)
	if err != nil {
		return "", err
	}
	err = config.DB.QueryRow(query, u.Email, u.Password, u.Name, u.LastName).Scan(&newUserId)
	if err != nil {
		return "", err
	}
	u.ID = newUserId
	return u.ID, nil
}

func (u UserLogin) Login() (loginResult, error) {
	var fetchedUser UserLogin
	var result loginResult
	var err error

	query := "SELECT user_id, email, password FROM users WHERE email = $1"
	err = config.DB.QueryRow(query, u.Email).Scan(&fetchedUser.ID, &fetchedUser.Email, &fetchedUser.Password)
	if err != nil {
		return loginResult{}, err
	}

	err = auth.ComparePassword(u.Password, fetchedUser.Password)
	if err != nil {
		return loginResult{}, err
	}

	result.UserId = fetchedUser.ID
	result.AccessToken, err = auth.GenerateAccessToken(fetchedUser.Email, fetchedUser.ID)
	if err != nil {
		return loginResult{}, err
	}
	result.RefreshToken = auth.GenerateRefreshToken()

	expDays, err := strconv.ParseInt(config.ServerCfg.RefreshExpiryDays, 10, 64)
	if err != nil {
		return loginResult{}, err
	}
	expDate := time.Now().UTC().AddDate(0, 0, int(expDays))
	result.RefreshTokenExpDate = expDate

	_, err = auth.PersistRefreshToken(utils.Hash256(result.RefreshToken), fetchedUser.ID, expDate)
	if err != nil {
		return loginResult{}, err
	}

	return result, nil
}
