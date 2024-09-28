package util_http_middleware

import (
	"github.com/gin-gonic/gin"

	util_error "github.com/Final-Project-Azie/e-commerce-be/util/error"
	util_http "github.com/Final-Project-Azie/e-commerce-be/util/http"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Content-Type", "application/json")
		c.Next()

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0]
		// if err can be casted to ClientError, then it is a client error
		if clientError, ok := err.Err.(*util_error.ClientError); ok {
			c.JSON(clientError.Code, util_http.Error{
				Message: clientError.Message,
			})
			return
		}

		if err.IsType(gin.ErrorTypeBind) {
			c.JSON(400, util_http.Error{
				Message: err.Err.Error(),
			})
			return
		}

		if err.IsType(gin.ErrorTypePrivate) {
			c.JSON(500, util_http.Error{
				Message: "Internal server error",
			})
			return
		}

		c.JSON(500, util_http.Error{
			Message: "Internal server error",
		})
	}
}
