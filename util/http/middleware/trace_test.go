package util_http_middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	util_http "backend-template/util/http"
	util_http_middleware "backend-template/util/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestTraceMiddleware(t *testing.T) {
	srv := util_http.NewHTTPServer(gin.TestMode)

	srv.Use(util_http_middleware.TraceIdAssignmentMiddleware())

	srv.GET("", func(ctx *gin.Context) {
		traceID, ok := ctx.Request.Context().Value(util_http.TraceString).(string)
		require.True(t, ok)
		require.NotEmpty(t, traceID)
	})

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	srv.ServeHTTP(res, req)
}
