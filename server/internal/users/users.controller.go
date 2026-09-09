package users

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sebasvelasco353/nummus/server/internal/auth"
	"github.com/sebasvelasco353/nummus/server/internal/config"
	"github.com/sebasvelasco353/nummus/server/internal/utils"
)

func signUp(context *gin.Context) {
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

	// TODO: secure=false — this needs to flip to true once we're serving over HTTPS
	context.SetCookie("nummus", result.RefreshToken, int(time.Until(result.RefreshTokenExpDate).Seconds()), "/refresh", "", false, true)

	context.JSON(200, gin.H{
		"message": "Success, user logged in with ID: " + result.UserId,
		"result": gin.H{
			"userID":      result.UserId,
			"accessToken": result.AccessToken,
		},
	})
}

// TODO: logout functionality
func logout(context *gin.Context) {
	fmt.Println("Logout the user!")
}

func refresh(context *gin.Context) {
	fmt.Println("Refresh the token")
	cookie, err := context.Cookie("nummus")
	if err != nil {
		context.JSON(401, gin.H{
			"error": err.Error(),
			"code":  "session_compromised",
		})
		return
	}
	cookie = utils.Hash256(cookie)
	result, err := auth.FetchRefreshToken(cookie)
	if err != nil {
		context.SetCookie("nummus", "", -1, "/refresh", "", false, true)
		context.JSON(401, gin.H{
			"error": err.Error(),
			"code":  "session_compromised",
		})
		return
	}

	// Token Validation before rotation
	if result.IsUsed {
		_, err = auth.RevokeAllRefreshTokens(result.UserID)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		context.SetCookie("nummus", "", -1, "/refresh", "", false, true)
		context.JSON(401, gin.H{
			"error": "session has been compromised, self destruct initiated.",
			"code":  "session_compromised",
		})
		return
	} else if result.ExpDate.Before(time.Now()) {
		context.SetCookie("nummus", "", -1, "/refresh", "", false, true)
		context.JSON(401, gin.H{
			"error": "session has expired, log in again.",
			"code":  "session_expired",
		})
		return
	} else {
		// TODO: wrap all rotation queries in a config.DB.Begin() / tx.Commit() / tx.Rollback() transaction so they succeed or fail together

		// Generates a new Access Token with claims
		claims := auth.AccessTokenClaims{
			Email:  result.Email,
			UserID: result.UserID,
		}
		expHours, err := strconv.ParseInt(config.ServerCfg.JWTExpiryHours, 10, 64)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expHours)))
		accessToken, err := auth.GenerateAccessToken(claims)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Generates a new Refresh Token
		expDays, err := strconv.ParseInt(config.ServerCfg.RefreshExpiryDays, 10, 64)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		refreshToken := auth.GenerateRefreshToken()
		refreshTokenExpDate := time.Now().UTC().AddDate(0, 0, int(expDays))
		_, err = auth.PersistRefreshToken(utils.Hash256(refreshToken), result.UserID, refreshTokenExpDate)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Revokes old refresh Token
		err = auth.RevokeSingleRefreshToken(result.RefreshTokenID)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.SetCookie("nummus", refreshToken, int(time.Until(refreshTokenExpDate).Seconds()), "/refresh", "", false, true)

		context.JSON(200, gin.H{
			"message": "Success, user refreshed their tokens",
			"result": gin.H{
				"userID":      result.UserID,
				"accessToken": accessToken,
			},
		})
	}
}

func getSelf(context *gin.Context) {
	context.JSON(200, gin.H{
		"userId": context.MustGet("userId"),
		"email":  context.MustGet("email"),
	})
}
