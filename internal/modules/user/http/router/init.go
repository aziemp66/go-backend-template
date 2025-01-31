package user_http_router

import (
	user_http "backend-template/internal/modules/user/http"
	util_http_middleware "backend-template/util/http/middleware"
	util_jwt "backend-template/util/jwt"

	"github.com/gin-gonic/gin"
)

func BindUserHttpRouter(router *gin.RouterGroup, handler user_http.UserHttpHandler, jwtManager util_jwt.JWTManager) {
	router.GET("/:id", handler.GetUserByID)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)

	userRoutes := router.Group(
		"/user",
		util_http_middleware.JWTAuthentication(jwtManager),
		util_http_middleware.JWTAuthorization(util_jwt.USER_ROLE),
	)
	userRoutes.PUT("", handler.UpdateProfile)
	userRoutes.POST("/change-password", handler.ChangePassword)
}
