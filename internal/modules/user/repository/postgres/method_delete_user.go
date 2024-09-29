package user_repository_postgres

import "context"

// DeleteUser removes a user from the database based on their ID.
// It returns an error if the delete operation fails.
func (userRepositoryPostgres *userRepositoryPostgres) DeleteUser(ctx context.Context, id string) (err error) {
	_, err = userRepositoryPostgres.db.ExecContext(ctx, deleteUserQuery, id)

	return err
}
