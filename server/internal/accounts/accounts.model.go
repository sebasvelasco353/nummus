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

func GetAccount(accountId string, userId string) (Account, error) {
	var account Account
	var err error

	query := "SELECT account_id, owner, bank, name, balance, currency, account_type FROM accounts WHERE account_id = $1 AND owner = $2"
	err = config.DB.QueryRow(query, accountId, userId).Scan(&account.AccountId, &account.Owner, &account.Bank, &account.Name, &account.Balance, &account.Currency, &account.AccountType)

	return account, err
}
