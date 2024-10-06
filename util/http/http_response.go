package util_http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SendResponseJson(ctx *gin.Context, message string, data any) {
	ctx.Header("Content-Type", "application/json")
	timestamp := time.Now().Format(time.RFC3339)

	ctx.JSON(http.StatusOK, Response{
		Metadata: Metadata{
			Message:   message,
			Timestamp: timestamp,
		},
		Data: data,
	})
}
