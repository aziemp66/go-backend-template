package user_repository_postgres

import (
	"context"
)

// Implements UserRepository
// TODO: Comment Here
func (userRepositoryPostgres *userRepositoryPostgres) GetUserByEmail(ctx context.Context, email string) (id string, err error) {
	panic("unimplemented")
}
