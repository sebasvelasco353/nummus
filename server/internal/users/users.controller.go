package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	// TODO: add validation for the user input
	// TODO: add logic to save the user to the database
	var user User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("User:", user)
	context.JSON(200, gin.H{
		"message": "Sign Up",
	})
}
