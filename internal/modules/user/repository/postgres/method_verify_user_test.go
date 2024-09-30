package user_repository_postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryVerifyUser(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repoMock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	reqEmail := "johndoe@gmail.com"

	t.Run("should verify user successfully", func(t *testing.T) {
		sqlMock.ExpectExec(verifyUserQuery).
			WithArgs(reqEmail).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repoMock.VerifyUser(context.Background(), reqEmail)
		assert.Nil(t, err)
	})

	t.Run("should return error when verify fails", func(t *testing.T) {
		sqlMock.ExpectExec(verifyUserQuery).
			WillReturnError(errors.New("db error"))

		err := repoMock.VerifyUser(context.Background(), "")
		assert.EqualError(t, err, "db error")
	})
}
