package user_repository_postgres

import "context"

// ChangePassword updates a user's password in the database based on their email.
// new password must be hashed before function and returns an error if the update fails.
func (userRepositoryPostgres *userRepositoryPostgres) ChangePassword(ctx context.Context, id string, hashedPassword string) (err error) {
	_, err = userRepositoryPostgres.db.ExecContext(ctx, changePasswordQuery, id, hashedPassword)

	return err
}
