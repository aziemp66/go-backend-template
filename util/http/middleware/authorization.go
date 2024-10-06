package util_http_middleware

import (
	util_error "backend-template/util/error"
	util_jwt "backend-template/util/jwt"
	"errors"

	"github.com/gin-gonic/gin"
)

func JWTAuthorization(roles ...util_jwt.ROLE) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("user_role")

		var isValid bool = false
		for _, v := range roles {
			if v == util_jwt.ROLE(userRole) {
				isValid = true
			}
		}

		if !isValid {
			c.Error(
				util_error.NewUnauthorized(errors.New("you dont have the privileges to access this route"),
					"you dont have the privileges to access this route",
				))
			c.Abort()
			return
		}

		c.Next()
	}
}
