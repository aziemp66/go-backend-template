package user_repository_postgres

import (
	"context"

	user_model "backend-template/internal/modules/user/model"
)

// GetUserByID retrieves a user from the database based on their ID.
// It returns the user details or an error if the query fails.
func (userRepositoryPostgres *userRepositoryPostgres) GetUserByID(ctx context.Context, id string) (res user_model.User, err error) {
	err = userRepositoryPostgres.db.GetContext(ctx, &res, getUserByID, id)

	if err != nil {
		return user_model.User{}, err
	}

	return res, nil
}
