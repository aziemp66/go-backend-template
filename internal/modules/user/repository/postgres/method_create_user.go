package user_repository_postgres

import (
	"context"

	user_model "github.com/Final-Project-Azie/e-commerce-be/internal/modules/user/model"
)

// Implements UserRepository
// TODO: Comment Here
func (userRepositoryPostgres *userRepositoryPostgres) CreateUser(ctx context.Context, name string, address string, email string, password string) (res user_model.User, err error) {
	panic("unimplemented")
}
