package user_service

import (
	"context"

	user_model "backend-template/internal/modules/user/model"
)

func (userService *userService) GetUserByEmail(ctx context.Context, email string) (res user_model.GetUserResponse, err error) {
	user, err := userService.userRepository.GetUserByEmail(ctx, email)
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
