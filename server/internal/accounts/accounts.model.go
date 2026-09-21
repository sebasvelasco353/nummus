package accounts

import (
	"github.com/sebasvelasco353/nummus/server/internal/config"
)

type Account struct {
	AccountId   string
	Owner       string
	Bank        string `binding:"required"`
	Name        string `binding:"required"`
	Balance     int
	Currency    string `binding:"required"`
	AccountType string `binding:"required,oneof=savings debit credit"`
}

func (a Account) CreateAccount() (string, error) {
	var err error
	var accountId string

	query := "INSERT INTO accounts (owner, bank, name, balance, currency, account_type) VALUES ($1, $2, $3, $4, $5, $6) returning account_id"
	err = config.DB.QueryRow(query, a.Owner, a.Bank, a.Name, a.Balance, a.Currency, a.AccountType).Scan(&accountId)
	if err != nil {
		return "", err
	}
	a.AccountId = accountId
	return a.AccountId, nil
}
