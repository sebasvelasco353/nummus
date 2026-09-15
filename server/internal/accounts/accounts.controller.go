package accounts

import "github.com/gin-gonic/gin"

func getAccounts(context *gin.Context) {
	context.JSON(200, gin.H{
		"success": true,
		"data":    nil,
	})
}
func getAccount(context *gin.Context) {
	context.JSON(200, gin.H{
		"success": true,
		"data":    nil,
	})
}

func createAccount(context *gin.Context) {
	context.JSON(200, gin.H{
		"success": true,
		"data":    nil,
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
