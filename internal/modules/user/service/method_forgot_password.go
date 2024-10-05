package user_service

import (
	util_error "backend-template/util/error"
	util_jwt "backend-template/util/jwt"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func (userService *userService) ForgotPassword(ctx context.Context, email string) error {
	user, err := userService.userRepository.GetUserByEmail(ctx, email)
	if err == sql.ErrNoRows {
		return util_error.NewNotFound(sql.ErrNoRows, fmt.Sprintf("User with the email of %s doesn't exist", email))
	} else if err != nil {
		return err
	}

	token, err := userService.jwtManager.GenerateAuthToken(user.ID, user.Name, util_jwt.USER_ROLE, 1*time.Hour)
	if err != nil {
		return err
	}

	err = userService.mailManager.SentResetPassword(token, email)
	if err != nil {
		return err
	}

	return nil
}
