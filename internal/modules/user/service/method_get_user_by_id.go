package user_service

import (
	"context"
	"database/sql"
	"fmt"

	user_model "backend-template/internal/modules/user/model"
	util_error "backend-template/util/error"
)

func (userService *userService) GetUserByID(ctx context.Context, id string) (res user_model.GetUserResponse, err error) {
	user, err := userService.userRepository.GetUserByID(ctx, id)
	if err == sql.ErrNoRows {
		return user_model.GetUserResponse{}, util_error.NewNotFound(err, fmt.Sprintf("User with the id of %s is not found", id))
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
