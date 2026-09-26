package transactions

import (
	"github.com/gin-gonic/gin"
	"github.com/sebasvelasco353/nummus/server/internal/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	var transactions = server.Group("/transactions")
	transactions.Use(middleware.AuthValidator())
	{
		transactions.GET("", getByUser)
		transactions.GET("/:id", getAccount)
		transactions.POST("", create)
		transactions.PUT("/:id", update)
		transactions.DELETE("/:id", delete)
	}
}
