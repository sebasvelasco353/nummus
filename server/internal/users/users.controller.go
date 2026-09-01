package users

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sebasvelasco353/nummus/server/internal/auth"
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

	/* TODO: domain="localhost" hardcoded — this works for local testing, but will fail when this deploys somewhere with a real domain. already have cfg.IsDevelopment() in config.go how can i implement this functionality? - Do we really need it? this will be self hosted :V
	 */
	/* TODO: secure=false — this needs to flip to true once we're serving over HTTPS
	 */

	context.SetCookie("nummus", result.RefreshToken, int(time.Until(result.RefreshTokenExpDate).Seconds()), "/refresh", "localhost", false, true)

	auth.ValidateAccessToken(result.AccessToken)

	context.JSON(200, gin.H{
		"message": "Success, user logged in with ID: " + result.UserId,
		"result": gin.H{
			"userID":      result.UserId,
			"accessToken": result.AccessToken,
		},
	})
}

func refresh(context *gin.Context) {
	fmt.Println("Refresh the token")
}
