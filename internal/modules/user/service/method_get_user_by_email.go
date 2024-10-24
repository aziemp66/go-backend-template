package user_service

import (
	"context"
	"database/sql"
	"fmt"

	user_model "backend-template/internal/modules/user/model"
	util_error "backend-template/util/error"
)

func (userService *userService) GetUserByEmail(ctx context.Context, email string) (res user_model.GetUserResponse, err error) {
	user, err := userService.userRepository.GetUserByEmail(ctx, email)
	if err == sql.ErrNoRows {
		return user_model.GetUserResponse{}, util_error.NewNotFound(err, fmt.Sprintf("User with the email of %s is not found", email))
	}
	if err != nil {
		return user_model.GetUserResponse{}, err
	}

	return user_model.GetUserResponse{
		ID:      user.ID,
		Email:   user.Email,
		Name:    user.Name,
		Address: user.Address,
	}, nil
}
