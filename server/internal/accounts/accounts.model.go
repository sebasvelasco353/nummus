package accounts

type Account struct {
	AccountId string
	Owner     string
	Bank      string `binding:"required"`
	Name      string `binding:"required"`
	Balance   int64  `binding:"required"`
	Currency  string `binding:"required"`
}
