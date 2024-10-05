package user_service

import (
	user_model "backend-template/internal/modules/user/model"
	mock_repository "backend-template/mock/repository"
	mock_util "backend-template/mock/util"
	util_error "backend-template/util/error"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)
	mailMock := mock_util.NewMockMailManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock, mailMock)

	id := "1"
	name := "John Doe"
	address := "123 Sesame Street"

	repoRes := user_model.User{
		ID:       id,
		Email:    "johndoe@gmail.com",
		Password: "secure_password",
		Name:     name,
		Address:  address,
	}

	t.Run("should update user successfully", func(t *testing.T) {
		repoMock.EXPECT().GetUserByID(gomock.Any(), id).Return(repoRes, nil)

		repoMock.EXPECT().UpdateUser(gomock.Any(), id, name, address).Return(nil)

		err := service.UpdateUser(context.Background(), id, name, address)

		require.NoError(t, err)
	})

	t.Run("should return error when update user fails", func(t *testing.T) {
		expectedErr := errors.New("failed to update user")

		repoMock.EXPECT().GetUserByID(gomock.Any(), "").Return(user_model.User{}, expectedErr)

		err := service.UpdateUser(context.Background(), "", "", "")

		require.Error(t, err)
		assert.EqualError(t, err, expectedErr.Error())
	})

	t.Run("should return client error when user is not found", func(t *testing.T) {
		expectedErr := util_error.NewNotFound(fmt.Errorf("user with the id of %s is not found", id), "User not found")
		repoMock.EXPECT().GetUserByID(gomock.Any(), id).Return(user_model.User{}, expectedErr)

		err := service.UpdateUser(context.Background(), id, "", "")

		require.Error(t, err)
		assert.EqualError(t, err, expectedErr.Error())
	})
}
