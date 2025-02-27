package util_http_middleware

import (
	"fmt"
	"time"

	util_logger "backend-template/util/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timer
		path := c.Request.URL.Path
		queryParam := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.BodySize = c.Writer.Size()
		if queryParam != "" {
			path = path + "?" + queryParam
		}
		param.Path = path

		if len(c.Errors) > 0 && param.StatusCode >= 500 && param.StatusCode < 600 {
			param.ErrorMessage = c.Errors[0].Error()
			util_logger.Error(
				c.Request.Context(),
				"Internal Server Error : "+param.ErrorMessage,
				zap.String("client_id", param.ClientIP),
				zap.String("http_method", param.Method),
				zap.Int("body_size", param.BodySize),
				zap.String("path", param.Path),
				zap.String("latency", param.Latency.String()),
				zap.Int("http_response_code", param.StatusCode),
			)
		} else {
			util_logger.Debug(
				c.Request.Context(),
				fmt.Sprintf("Request %s [%s]", path, param.Method),
				zap.String("client_id", param.ClientIP),
				zap.String("http_method", param.Method),
				zap.Int("body_size", param.BodySize),
				zap.String("path", path),
				zap.String("latency", param.Latency.String()),
				zap.Int("http_response_code", param.StatusCode),
			)
		}
	}
}
