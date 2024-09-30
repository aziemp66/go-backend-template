package user_service

import (
	"context"

	user_model "github.com/Final-Project-Azie/e-commerce-be/internal/modules/user/model"
)

type UserService interface {
	Register(ctx context.Context, email, password, name, address string) (id string, err error)
	GetUserByID(ctx context.Context, id string) (res user_model.GetUserResponse, err error)
	GetUserByEmail(ctx context.Context, email string) (res user_model.GetUserResponse, err error)
	Login(ctx context.Context, email, password string) (token string, err error)
	ForgotPassword(ctx context.Context, email string)
	ResetPassword(ctx context.Context, token, email, newPassword string) (err error)
	VerifyUser(ctx context.Context, email string) (err error)
}
