package accounts

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func handleGetAll(context *gin.Context) {
	ownerId := context.MustGet("userId").(string)

	result, err := GetAllByOwner(ownerId)
	if err != nil {
		handleErrors(context, err)
		return
	}

	context.JSON(200, gin.H{
		"success": true,
		"data":    result,
	})
}

func handleGetOne(context *gin.Context) {
	accountId := context.Param("id")
	ownerId := context.MustGet("userId").(string)

	result, err := GetOneByOwner(accountId, ownerId)
	if err != nil {
		handleErrors(context, err)
		return
	}

	context.JSON(200, gin.H{
		"success": true,
		"data":    result,
	})
}

func handleCreate(context *gin.Context) {
	var account Account

	if err := context.ShouldBindBodyWithJSON(&account); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ownerId := context.MustGet("userId")
	account.Owner = ownerId.(string)

	resultId, err := account.insert()
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

func handleUpdate(context *gin.Context) {
	var newData UpdateAccountData
	var err error

	ownerId := context.MustGet("userId").(string)
	accountId := context.Param("id")

	if err = context.ShouldBindBodyWithJSON(&newData); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedAccount, err := updateDetails(accountId, ownerId, newData)
	if err != nil {
		handleErrors(context, err)
		return
	}

	context.JSON(200, gin.H{
		"success": true,
		"data":    updatedAccount,
	})
}

func handleDelete(context *gin.Context) {
	accountId := context.Param("id")
	ownerId := context.MustGet("userId").(string)

	affectedRows, err := deleteOne(accountId, ownerId)
	if err != nil {
		handleErrors(context, err)
		return
	}
	if affectedRows == 0 {
		handleErrors(context, sql.ErrNoRows)
		return
	}
	context.Status(204)
}
