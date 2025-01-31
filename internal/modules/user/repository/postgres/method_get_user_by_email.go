package user_repository_postgres

import (
	"context"

	user_model "backend-template/internal/modules/user/model"
)

// GetUserByEmail retrieves a user from the database based on their email.
// It returns the user details or an error if the query fails.
func (userRepositoryPostgres *userRepositoryPostgres) GetUserByEmail(ctx context.Context, email string) (res user_model.User, err error) {
	err = userRepositoryPostgres.db.GetContext(ctx, &res, getUserByEmail, email)

	if err != nil {
		return user_model.User{}, err
	}

	return res, nil
}
