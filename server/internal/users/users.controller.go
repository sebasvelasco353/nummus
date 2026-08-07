package users

import (
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	// TODO: add validation for the user input
	var user User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	result, err := user.SignUp()
	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"message": "Success, user created with ID: " + result,
	})
}

func login(context *gin.Context) {
	// TODO: Add JWT token generation functionality
	var user UserLogin

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	result, err := user.Login()
	if err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(200, gin.H{
		"message": "Success, user logged in with ID: " + result,
	})
}
