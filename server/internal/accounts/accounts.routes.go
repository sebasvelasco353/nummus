package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/sebasvelasco353/nummus/server/internal/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	var accounts = server.Group("/accounts")
	accounts.Use(middleware.AuthValidator())
	{
		accounts.GET("", getAccounts)
		accounts.GET("/:id", getAccount)
		accounts.POST("", createAccount)
		accounts.PUT("/:id", updateAccount)
		accounts.DELETE("/:id", deleteAccount)
	}
}
