package user_http

import (
	user_service "backend-template/internal/modules/user/service"
	util_http_middleware "backend-template/util/http/middleware"
	util_jwt "backend-template/util/jwt"

	"github.com/gin-gonic/gin"
)

func NewUserHttpHandler(router *gin.RouterGroup, userService user_service.UserService, jwtManager util_jwt.JWTManager) {
	handler := userHttpHandler{
		userService: userService,
		jwtManager:  jwtManager,
	}

	router.GET("/:id", handler.GetUserByID)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.POST("/forgot-password", handler.ForgotPassword)
	router.POST("/reset-password/:token", handler.ResetPassword)
	router.POST("/verify-user/:token", handler.VerifyUser)

	userRoutes := router.Group(
		"/user",
		util_http_middleware.JWTAuthentication(jwtManager),
		util_http_middleware.JWTAuthorization(util_jwt.USER_ROLE),
	)
	userRoutes.PUT("", handler.UpdateProfile)
	userRoutes.POST("/change-password", handler.ChangePassword)
}
