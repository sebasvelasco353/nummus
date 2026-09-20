package accounts

import (
	"github.com/sebasvelasco353/nummus/server/internal/config"
)

type Account struct {
	AccountId string
	Owner     string
	Bank      string `binding:"required"`
	Name      string `binding:"required"`
	Balance   int
	Currency  string `binding:"required"`
}

func (a Account) CreateAccount() (string, error) {
	var err error
	var accountId string

	query := "INSERT INTO accounts (owner, bank, name, balance, currency) VALUES ($1, $2, $3, $4, $5) returning account_id"
	err = config.DB.QueryRow(query, a.Owner, a.Bank, a.Name, a.Balance, a.Currency).Scan(&accountId)
	if err != nil {
		return "", err
	}
	a.AccountId = accountId
	return a.AccountId, nil
}
