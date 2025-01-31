package user_service

import (
	util_error "backend-template/util/error"
	util_jwt "backend-template/util/jwt"
	"context"
	"database/sql"
	"time"
)

func (userService *userService) Login(ctx context.Context, email string, password string) (token string, err error) {
	user, err := userService.userRepository.GetUserByEmail(ctx, email)
	if err == sql.ErrNoRows {
		return "", util_error.NewUnauthorized(err, "Wrong Email or Password")
	} else if err != nil {
		return "", err
	}

	if err := userService.passwordManager.CheckPasswordHash(password, user.Password); err != nil {
		return "", err
	}

	token, err = userService.jwtManager.GenerateAuthToken(user.ID, user.Name, util_jwt.USER_ROLE, 24*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}
