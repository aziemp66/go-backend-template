package user_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	user_model "backend-template/internal/modules/user/model"
	mock_repository "backend-template/mock/repository"
	mock_util "backend-template/mock/util"
	util_error "backend-template/util/error"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock)

	idReq := "123"

	userRes := user_model.GetUserResponse{
		ID:      idReq,
		Name:    "John Smith",
		Address: "Sesame Street No.5",
		Email:   "johnsmith123@gmail.com",
	}

	t.Run("should get user by id", func(t *testing.T) {
		repoRes := user_model.User{
			ID:         idReq,
			Email:      userRes.Email,
			Password:   "secured_password",
			Name:       userRes.Name,
			Address:    userRes.Address,
			IsVerified: true,
		}

		repoMock.EXPECT().GetUserByID(gomock.Any(), idReq).
			Return(repoRes, nil)

		res, err := service.GetUserByID(context.Background(), idReq)

		require.Nil(t, err)
		assert.Equal(t, userRes, res)
	})

	t.Run("should return client error when theres no user equal to id requirement", func(t *testing.T) {
		expectedErr := util_error.NewNotFound(sql.ErrNoRows, fmt.Sprintf("User with the id of %s is not found", idReq))

		repoMock.EXPECT().GetUserByID(gomock.Any(), idReq).
			Return(user_model.User{}, sql.ErrNoRows)

		_, err := service.GetUserByID(context.Background(), idReq)

		require.Error(t, err)
		assert.EqualError(t, expectedErr, err.Error())
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("error from db")

		repoMock.EXPECT().GetUserByID(gomock.Any(), idReq).
			Return(user_model.User{}, expectedErr)

		_, err := service.GetUserByID(context.Background(), idReq)
		require.Error(t, err)
		assert.EqualError(t, expectedErr, err.Error())
	})
}
