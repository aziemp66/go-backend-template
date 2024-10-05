package user_service

import (
	util_error "backend-template/util/error"
	util_jwt "backend-template/util/jwt"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (userService *userService) Register(ctx context.Context, email string, password string, name string, address string) (id string, err error) {
	_, err = userService.userRepository.GetUserByEmail(ctx, email)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	} else if err == nil {
		return "", util_error.NewBadRequest(fmt.Errorf("%s is already registered", email), fmt.Sprintf("Email %s is already used", email))
	}

	if err := userService.passwordManager.PasswordValidation(password); err != nil {
		return "", err
	}

	hashedPassword, err := userService.passwordManager.HashPassword(password)
	if err != nil {
		return "", err
	}

	id, err = userService.userRepository.CreateUser(ctx, email, hashedPassword, name, address)
	if err != nil {
		return "", err
	}

	token, err := userService.jwtManager.GenerateAuthToken(id, name, util_jwt.USER_ROLE, 1*time.Hour)
	if err != nil {
		return "", err
	}

	err = userService.mailManager.SentVerifyEmail(token, email)
	if err != nil {
		return "", err
	}

	return id, nil
}
