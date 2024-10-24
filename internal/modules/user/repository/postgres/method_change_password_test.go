package user_repository_postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryChangePassword(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repoMock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	reqID := "123"
	reqHashed := "hashedPassword"
	if err != nil {
		panic(err)
	}

	t.Run("should run execute update query", func(t *testing.T) {
		sqlMock.ExpectExec(changePasswordQuery).
			WithArgs(reqID, reqHashed).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err = repoMock.ChangePassword(context.Background(), reqID, reqHashed)
		assert.Nil(t, err)
	})

	t.Run("should return error when failed update query", func(t *testing.T) {
		rowErr := errors.New("db error")
		sqlMock.ExpectExec(changePasswordQuery).
			WillReturnError(rowErr)

		err := repoMock.ChangePassword(context.Background(), "", "")

		assert.EqualError(t, err, rowErr.Error())
	})
}
