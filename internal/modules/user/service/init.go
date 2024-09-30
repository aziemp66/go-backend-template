package user_service

import (
	user_repository "github.com/Final-Project-Azie/e-commerce-be/internal/modules/user/repository"
	util_jwt "github.com/Final-Project-Azie/e-commerce-be/util/jwt"
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
