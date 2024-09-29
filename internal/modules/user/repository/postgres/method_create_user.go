package user_repository_postgres

import (
	"context"
)

// CreateUser inserts a new user into the database with the provided details.
// It returns the newly created user's ID or an error if the operation fails.
func (userRepositoryPostgres *userRepositoryPostgres) CreateUser(ctx context.Context, email string, hashedPassword, name string, address string) (id string, err error) {
	row := userRepositoryPostgres.db.QueryRowxContext(ctx, createUserQuery, email, hashedPassword, name, address)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
