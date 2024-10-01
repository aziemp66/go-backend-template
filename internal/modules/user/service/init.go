package user_service

import (
	user_repository "backend-template/internal/modules/user/repository"
	util_jwt "backend-template/util/jwt"
)

type userService struct {
	userRepository user_repository.UserRepository
	jwtManager     *util_jwt.JWTManager
}

func NewUserService(userRepository user_repository.UserRepository, jwtManager *util_jwt.JWTManager) UserService {
	return &userService{
		userRepository: userRepository,
		jwtManager:     jwtManager,
	}
}
