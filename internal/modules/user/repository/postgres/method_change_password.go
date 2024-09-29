package user_repository_postgres

import "context"

// Implements UserRepository
// TODO: Comment Here
func (userRepositoryPostgres *userRepositoryPostgres) ChangePassword(ctx context.Context, email string, newPassword string) (err error) {
	panic("unimplemented")
}
