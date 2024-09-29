package user_repository_postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryDeleteUser(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repoMock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	userID := "1"

	t.Run("should delete user successfully", func(t *testing.T) {
		sqlMock.ExpectExec(deleteUserQuery).
			WithArgs(userID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repoMock.DeleteUser(context.Background(), userID)
		assert.Nil(t, err)
	})

	t.Run("should return error when delete fails", func(t *testing.T) {
		sqlMock.ExpectExec(deleteUserQuery).
			WithArgs(userID).
			WillReturnError(errors.New("db error"))

		err := repoMock.DeleteUser(context.Background(), userID)
		assert.EqualError(t, err, "db error")
	})
}
