package accounts

import (
	"github.com/gin-gonic/gin"
)

func getAccounts(context *gin.Context) {
	context.JSON(200, gin.H{
		"success": true,
		"data":    nil,
	})
}
func getAccount(context *gin.Context) {
	accountId := context.Param("id")
	ownerId := context.MustGet("userId").(string)

	result, err := GetAccount(accountId, ownerId)
	if err != nil {
		handleErrors(context, err)
		return
	}

	context.JSON(200, gin.H{
		"success": true,
		"data":    result,
	})
}

func createAccount(context *gin.Context) {
	var account Account

	if err := context.ShouldBindBodyWithJSON(&account); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ownerId := context.MustGet("userId")
	account.Owner = ownerId.(string)

	resultId, err := account.CreateAccount()
	if err != nil {
		handleErrors(context, err)
		return
	}

	account.AccountId = resultId
	context.JSON(201, gin.H{
		"success": true,
		"data":    account,
	})
}

func updateAccount(context *gin.Context) {
	context.JSON(200, gin.H{
		"success": true,
		"data":    nil,
	})
}

func deleteAccount(context *gin.Context) {
	context.JSON(200, gin.H{
		"success": true,
		"data":    nil,
	})
}
