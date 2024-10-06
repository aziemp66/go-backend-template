package user_http

import (
	user_service "backend-template/internal/modules/user/service"
	util_jwt "backend-template/util/jwt"
)

type userHttpHandler struct {
	userService user_service.UserService
	jwtManager  util_jwt.JWTManager
}
