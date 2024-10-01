package user_repository_postgres

import (
	"context"
	"errors"
	"testing"

	user_model "backend-template/internal/modules/user/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryGetUserByEmail(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repoMock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	reqEmail := "johndoe@example.com"
	expectedUser := user_model.User{
		ID:       "1",
		Email:    reqEmail,
		Password: "securepassword",
		Name:     "John Doe",
		Address:  "123 Main St",
	}

	t.Run("should return user by email", func(t *testing.T) {
		sqlMock.ExpectQuery(getUserByEmail).
			WithArgs(reqEmail).
			WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "name", "address"}).
				AddRow(expectedUser.ID, expectedUser.Email, expectedUser.Password, expectedUser.Name, expectedUser.Address))

		user, err := repoMock.GetUserByEmail(context.Background(), reqEmail)
		assert.Nil(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("should return error when query fails", func(t *testing.T) {
		sqlMock.ExpectQuery(getUserByEmail).
			WithArgs(reqEmail).
			WillReturnError(errors.New("db error"))

		user, err := repoMock.GetUserByEmail(context.Background(), reqEmail)
		assert.Empty(t, user)
		assert.EqualError(t, err, "db error")
	})
}
