package user_repository_postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryUpdateUser(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repoMock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	userID := "1"
	reqName := "John Doe Updated"
	reqAddress := "456 New St"

	t.Run("should update user successfully", func(t *testing.T) {
		sqlMock.ExpectExec(updateUserQuery).
			WithArgs(userID, reqName, reqAddress).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repoMock.UpdateUser(context.Background(), userID, reqName, reqAddress)
		assert.Nil(t, err)
	})

	t.Run("should return error when update fails", func(t *testing.T) {
		sqlMock.ExpectExec(updateUserQuery).
			WithArgs(userID, reqName, reqAddress).
			WillReturnError(errors.New("db error"))

		err := repoMock.UpdateUser(context.Background(), userID, reqName, reqAddress)
		assert.EqualError(t, err, "db error")
	})
}
