package user_service

import "context"

// Implements UserService
// TODO: Comment Here
func (userService *userService) ResetPassword(ctx context.Context, token string, email string, newPassword string) (err error) {
	panic("implement me")
}
