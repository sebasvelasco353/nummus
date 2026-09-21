package accounts

import (
	"database/sql"
	"errors"
	"log/slog"
	"os"

	"github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func handleErrors(context *gin.Context, err error) {

	if errors.Is(err, sql.ErrNoRows) {
		context.JSON(404, gin.H{
			"error": "account not found",
		})
		return
	}

	var errCode *pq.Error
	if errors.As(err, &errCode) {
		if errCode.Code.Name() == "foreign_key_violation" {
			context.JSON(401, gin.H{
				"error": "no valid session found",
			})
			return
		}
		if errCode.Code.Name() == "invalid_text_representation" {
			context.JSON(400, gin.H{
				"error": "malformed id found",
			})
			return
		}
	}

	logger.Error("Request failed", "method", context.Request.Method, "path", context.FullPath(), "error", err.Error())
	context.JSON(500, gin.H{
		"error": "there was an unexpected error",
	})
}
