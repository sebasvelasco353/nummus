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
type UpdateAccountData struct {
	Name        string `binding:"required"`
	Currency    string `binding:"required"`
	AccountType string `binding:"required,oneof=savings debit credit"`
}

func (a Account) insert() (string, error) {
	var err error
	var accountId string

	query := "INSERT INTO accounts (owner, bank, name, balance, currency, account_type) VALUES ($1, $2, $3, $4, $5, $6) returning account_id"
	err = config.DB.QueryRow(query, a.Owner, a.Bank, a.Name, a.Balance, a.Currency, a.AccountType).Scan(&accountId)
	if err != nil {
		return "", err
	}
	return accountId, nil
}

func updateDetails(accountId string, owner string, newData UpdateAccountData) (Account, error) {
	var err error
	var updatedAccount Account

	query := "UPDATE accounts SET name = $1, currency = $2, account_type = $3 WHERE owner = $4 AND account_id = $5 RETURNING account_id, owner, bank, name, balance, currency, account_type"
	err = config.DB.QueryRow(query, newData.Name, newData.Currency, newData.AccountType, owner, accountId).Scan(&updatedAccount.AccountId, &updatedAccount.Owner, &updatedAccount.Bank, &updatedAccount.Name, &updatedAccount.Balance, &updatedAccount.Currency, &updatedAccount.AccountType)
	if err != nil {
		return Account{}, err
	}
	return updatedAccount, nil
}

func deleteOne(accountId string, owner string) (int64, error) {
	query := "DELETE FROM accounts WHERE account_id = $1 AND owner = $2"
	res, err := config.DB.Exec(query, accountId, owner)
	if err != nil {
		return 0, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetOneByOwner(accountId string, owner string) (Account, error) {
	var account Account
	var err error

	query := "SELECT account_id, owner, bank, name, balance, currency, account_type FROM accounts WHERE account_id = $1 AND owner = $2"
	err = config.DB.QueryRow(query, accountId, owner).Scan(&account.AccountId, &account.Owner, &account.Bank, &account.Name, &account.Balance, &account.Currency, &account.AccountType)

	return account, err
}

func GetAllByOwner(owner string) ([]Account, error) {
	var accounts []Account

	query := "SELECT account_id, owner, bank, name, balance, currency, account_type FROM accounts WHERE owner = $1"
	rows, err := config.DB.Query(query, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.AccountId, &account.Owner, &account.Bank, &account.Name, &account.Balance, &account.Currency, &account.AccountType); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(accounts) == 0 {
		accounts = make([]Account, 0)
	}
	return accounts, nil
}
