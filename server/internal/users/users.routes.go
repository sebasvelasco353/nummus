package users

import (
	"github.com/gin-gonic/gin"
	"github.com/sebasvelasco353/nummus/server/internal/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	private := server.Group("/user")
	private.Use(middleware.AuthValidator())
	{
		private.GET("/me", getSelf)
	}
	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.POST("/refresh", refresh)
}
