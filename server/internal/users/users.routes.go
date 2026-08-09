package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.POST("/refresh", refresh)
}
