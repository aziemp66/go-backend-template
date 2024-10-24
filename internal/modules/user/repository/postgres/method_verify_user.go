package user_repository_postgres

import "context"

func (userRepositoryPostgres *userRepositoryPostgres) VerifyUser(ctx context.Context, id string) (err error) {
	_, err = userRepositoryPostgres.db.ExecContext(ctx, verifyUserQuery, id)
	if err != nil {
		return err
	}

	return nil
}
