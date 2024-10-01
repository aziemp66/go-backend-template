package user_service

import "context"

func (userService *userService) ChangePassword(ctx context.Context, email, oldPassword, newPassword string) (err error) {
	panic("implement me")
}
