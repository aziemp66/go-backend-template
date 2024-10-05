package user_service

import (
	user_repository "backend-template/internal/modules/user/repository"
	util_jwt "backend-template/util/jwt"
	util_mail "backend-template/util/mail"
	util_password "backend-template/util/password"
)

type userService struct {
	userRepository  user_repository.UserRepository
	jwtManager      util_jwt.JWTManager
	passwordManager util_password.PasswordManager
	mailManager     util_mail.MailManager
}

func NewUserService(userRepository user_repository.UserRepository, jwtManager util_jwt.JWTManager, passwordManager util_password.PasswordManager, mailManager util_mail.MailManager) UserService {
	return &userService{
		userRepository:  userRepository,
		jwtManager:      jwtManager,
		passwordManager: passwordManager,
		mailManager:     mailManager,
	}
}
