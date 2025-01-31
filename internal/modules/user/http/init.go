package user_http

import (
	user_service "backend-template/internal/modules/user/service"
	util_jwt "backend-template/util/jwt"
)

func NewUserHttpHandler(userService user_service.UserService, jwtManager util_jwt.JWTManager) UserHttpHandler {
	return &userHttpHandler{
		userService: userService,
		jwtManager:  jwtManager,
	}
}
