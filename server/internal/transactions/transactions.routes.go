package transactions

import (
	"github.com/gin-gonic/gin"
	"github.com/sebasvelasco353/nummus/server/internal/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	var transactions = server.Group("/transactions")
	transactions.Use(middleware.AuthValidator())
	{
		transactions.GET("", handleGetAll)
		transactions.GET("/:id", handleGetOne)
		transactions.POST("", handleCreate)
		transactions.PUT("/:id", handleUpdate)
		transactions.DELETE("/:id", handleDelete)
	}
}
