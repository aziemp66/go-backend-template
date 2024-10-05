package user_service

import (
	util_error "backend-template/util/error"
	"context"
	"errors"
)

func (userService *userService) VerifyUser(ctx context.Context, email string) (err error) {
	user, err := userService.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user.IsVerified {
		return util_error.NewBadRequest(errors.New("user is already verified"), "User is already verified")
	}

	err = userService.userRepository.VerifyUser(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
