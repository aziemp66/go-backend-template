package util_http_middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	util_error "backend-template/util/error"
	util_http "backend-template/util/http"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) < 1 {
			return
		}
		c.Writer.Header().Add("Content-Type", "application/json")

		err := c.Errors[0]
		// if err can be casted to ClientError, then it is a client error
		if clientError, ok := err.Err.(*util_error.ClientError); ok {
			util_http.SendErrorResponseJson(c, clientError.Message, clientError.Code)
			return
		}

		if err.IsType(gin.ErrorTypeBind) {
			c.JSON(400, util_http.Error{
				Message: err.Err.Error(),
			})
			util_http.SendErrorResponseJson(c, err.Err.Error(), http.StatusBadRequest)
			return
		}

		if err.IsType(gin.ErrorTypePrivate) {
			util_http.SendErrorResponseJson(c, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		util_http.SendErrorResponseJson(c, "Internal Server Error", http.StatusInternalServerError)
	}
}
