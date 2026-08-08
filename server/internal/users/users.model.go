package users

import (
	"fmt"

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

func (u User) SignUp() (string, error) {
	var err error
	var newUserId string

	query := "INSERT INTO users (email, password, name, last_name) VALUES ($1, $2, $3, $4) RETURNING user_id"
	u.Password, err = utils.HashPassword(u.Password)
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

func (u UserLogin) Login() (string, error) {
	var fetchedUser UserLogin
	var err error

	query := "SELECT user_id, email, password FROM users WHERE email = $1"
	err = config.DB.QueryRow(query, u.Email).Scan(&fetchedUser.ID, &fetchedUser.Email, &fetchedUser.Password)
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(u.Password, fetchedUser.Password)
	if err != nil {
		return "", err
	}

	newToken, err := utils.GenerateToken(fetchedUser.Email, fetchedUser.ID)
	if err != nil {
		return "", err
	}
	fmt.Println(newToken)

	return fetchedUser.ID, nil
}
