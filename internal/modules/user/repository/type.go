package user_repository

import (
	"context"

	user_model "backend-template/internal/modules/user/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email, hashedPassword, name, address string) (id string, err error)
	GetUserByID(ctx context.Context, id string) (res user_model.User, err error)
	GetUserByEmail(ctx context.Context, email string) (res user_model.User, err error)
	ChangePassword(ctx context.Context, email, hashedPassword string) (err error)
	UpdateUser(ctx context.Context, id string, name, address string) (err error)
	DeleteUser(ctx context.Context, id string) (err error)
	VerifyUser(ctx context.Context, email string) (err error)
}
