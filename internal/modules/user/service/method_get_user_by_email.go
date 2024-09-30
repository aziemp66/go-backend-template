package user_service

import (
	"context"

	user_model "github.com/Final-Project-Azie/e-commerce-be/internal/modules/user/model"
)

// Implements UserService
// TODO: Comment Here
func (userService *userService) GetUserByEmail(ctx context.Context, email string) (res user_model.GetUserResponse, err error) {
	panic("implement me")
}
