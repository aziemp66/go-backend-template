package user_repository_postgres

import "context"

func (userRepositoryPostgres *userRepositoryPostgres) VerifyUser(ctx context.Context, email string) (err error) {
	_, err = userRepositoryPostgres.db.ExecContext(ctx, verifyUserQuery, email)
	if err != nil {
		return err
	}

	return nil
}
