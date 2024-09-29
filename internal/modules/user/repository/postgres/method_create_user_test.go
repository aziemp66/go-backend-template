package user_repository_postgres

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreateUser(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repoMock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	reqEmail := "johndoe@example.com"
	reqPassword := "securepassword"
	reqName := "John Doe"
	reqAddress := "123 Main St"
	returnedID := "1"

	t.Run("should insert user and return ID", func(t *testing.T) {
		sqlMock.ExpectQuery(createUserQuery).
			WithArgs(reqEmail, reqPassword, reqName, reqAddress).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(returnedID))

		id, err := repoMock.CreateUser(context.Background(), reqEmail, reqPassword, reqName, reqAddress)
		assert.Nil(t, err)
		assert.Equal(t, returnedID, id)
	})

	t.Run("should return error when insert fails", func(t *testing.T) {
		sqlMock.ExpectQuery(createUserQuery).
			WillReturnError(errors.New("db error"))

		id, err := repoMock.CreateUser(context.Background(), reqName, reqAddress, reqEmail, reqPassword)
		assert.Empty(t, id)
		assert.EqualError(t, err, "db error")
	})
}
