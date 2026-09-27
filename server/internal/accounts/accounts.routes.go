package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/sebasvelasco353/nummus/server/internal/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	var accounts = server.Group("/accounts")
	accounts.Use(middleware.AuthValidator())
	{
		accounts.GET("", handleGetAll)
		accounts.GET("/:id", handleGetOne)
		accounts.POST("", handleCreate)
		accounts.PUT("/:id", handleUpdate)
		accounts.DELETE("/:id", handleDelete)
	}
}
