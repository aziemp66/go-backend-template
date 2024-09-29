package user_repository_postgres

import (
	"context"

	user_model "github.com/Final-Project-Azie/e-commerce-be/internal/modules/user/model"
)

// GetUserByEmail retrieves a user from the database based on their email.
// It returns the user details or an error if the query fails.
func (userRepositoryPostgres *userRepositoryPostgres) GetUserByEmail(ctx context.Context, email string) (res user_model.User, err error) {
	row := userRepositoryPostgres.db.QueryRowxContext(ctx, getUserByEmail, email)

	if err = row.StructScan(&res); err != nil {
		return user_model.User{}, err
	}

	return res, nil
}
