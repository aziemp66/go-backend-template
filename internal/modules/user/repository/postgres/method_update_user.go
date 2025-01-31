package user_repository_postgres

import "context"

// UpdateUser updates a user's name and address in the database based on their ID.
// It returns an error if the update fails.
func (userRepositoryPostgres *userRepositoryPostgres) UpdateUser(ctx context.Context, id string, name string, address string) (err error) {
	_, err = userRepositoryPostgres.db.ExecContext(ctx, updateUserQuery, id, name, address)

	if err != nil {
		return err
	}

	return nil
}
