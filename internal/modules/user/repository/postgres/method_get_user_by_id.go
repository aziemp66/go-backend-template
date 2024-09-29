package user_repository_postgres

import (
	"context"

	user_model "github.com/Final-Project-Azie/e-commerce-be/internal/modules/user/model"
)

// Implements UserRepository
// TODO: Comment Here
func (userRepositoryPostgres *userRepositoryPostgres) GetUserByID(ctx context.Context, id string) (res user_model.User, err error) {
	panic("unimplemented")
}
