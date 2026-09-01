package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sebasvelasco353/nummus/server/internal/auth"
)

func AuthValidator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.GetHeader("Authorization")
		accessToken = strings.TrimPrefix(accessToken, "Bearer ")

		if accessToken == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "no access token in header"})
			return
		}

		claims, err := auth.ValidateAccessToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "access token not valid"})
			return
		}
		ctx.Set("userId", claims.UserID)
		ctx.Set("email", claims.Email)

		ctx.Next()
	}
}
