package util_http_middleware

import (
	"context"

	util_http "backend-template/util/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceIdAssignmentMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceContext := context.WithValue(ctx.Request.Context(), util_http.TraceString, uuid.NewString())
		httpReq := ctx.Request.WithContext(traceContext)
		ctx.Request = httpReq
	}
}
