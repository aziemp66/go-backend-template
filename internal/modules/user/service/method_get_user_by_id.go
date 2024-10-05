package user_service

import (
	"context"

	user_model "backend-template/internal/modules/user/model"
)

func (userService *userService) GetUserByID(ctx context.Context, id string) (res user_model.GetUserResponse, err error) {
	user, err := userService.userRepository.GetUserByID(ctx, id)
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
