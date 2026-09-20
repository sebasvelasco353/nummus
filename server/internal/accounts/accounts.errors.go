package accounts

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func handleErrors(context *gin.Context, err error) {
	logger.Error("Request failed", "method", context.Request.Method, "path", context.FullPath(), "error", err.Error())

	context.JSON(500, gin.H{
		"error": "there was an unexpected error",
	})
}
