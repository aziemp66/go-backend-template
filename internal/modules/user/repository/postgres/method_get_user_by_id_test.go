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

func TestUserRepositoryGetUserByID(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repoMock := NewUserRepositoryPostgres(sqlx.NewDb(db, "sqlmock"))

	userID := "1"
	expectedUser := user_model.User{
		ID:       userID,
		Email:    "johndoe@example.com",
		Password: "securepassword",
		Name:     "John Doe",
		Address:  "123 Main St",
	}

	t.Run("should return user by ID", func(t *testing.T) {
		sqlMock.ExpectQuery(getUserByID).
			WithArgs(userID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "name", "address"}).
				AddRow(expectedUser.ID, expectedUser.Email, expectedUser.Password, expectedUser.Name, expectedUser.Address))

		user, err := repoMock.GetUserByID(context.Background(), userID)
		assert.Nil(t, err)
		assert.Equal(t, expectedUser, user)
	})

	t.Run("should return error when query fails", func(t *testing.T) {
		sqlMock.ExpectQuery(getUserByID).
			WithArgs(userID).
			WillReturnError(errors.New("db error"))

		user, err := repoMock.GetUserByID(context.Background(), userID)
		assert.Empty(t, user)
		assert.EqualError(t, err, "db error")
	})
}
