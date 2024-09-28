package user_repository

import (
	"context"

	user_model "github.com/Final-Project-Azie/e-commerce-be/internal/modules/user/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, name, address, email, password string) (user_model.User, error)
	GetUserByID(ctx context.Context, id string) (user_model.User, error)
	GetUserByEmail(ctx context.Context, email string) (user_model.User, error)
	ChangePassword(ctx context.Context, email, oldPassword, newPassword string) error
	UpdateUser(ctx context.Context, id string, name, address string) error
	DeleteUser(ctx context.Context, id string) error
}
