package user_repository_postgres

import (
	"context"

	user_model "github.com/Final-Project-Azie/e-commerce-be/internal/modules/user/model"
)

// GetUserByID retrieves a user from the database based on their ID.
// It returns the user details or an error if the query fails.
func (userRepositoryPostgres *userRepositoryPostgres) GetUserByID(ctx context.Context, id string) (res user_model.User, err error) {
	row := userRepositoryPostgres.db.QueryRowxContext(ctx, getUserByID, id)

	if err = row.StructScan(&res); err != nil {
		return user_model.User{}, err
	}

	return res, nil
}
