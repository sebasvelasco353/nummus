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
	fmt.Println("New user ID:", newUserId)
	u.ID = newUserId
	fmt.Println("Signing up user:", u.ID)
	return u.ID, nil
}
