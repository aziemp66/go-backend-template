package user_service

import (
	user_model "backend-template/internal/modules/user/model"
	mock_repository "backend-template/mock/repository"
	mock_util "backend-template/mock/util"
	util_error "backend-template/util/error"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceVerifyUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)
	mailMock := mock_util.NewMockMailManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock, mailMock)

	userID := "123"

	repoRes := user_model.User{
		ID:         userID,
		Email:      "john@gmail.com",
		Password:   "secure_password",
		Name:       "Joe",
		Address:    "Sesame Street",
		IsVerified: false,
	}

	t.Run("should verify user successfully", func(t *testing.T) {
		repoMock.EXPECT().GetUserByID(gomock.Any(), userID).Return(repoRes, nil)

		repoMock.EXPECT().VerifyUser(gomock.Any(), userID).
			Return(nil)

		err := service.VerifyUser(context.Background(), userID)

		require.NoError(t, err)
		assert.Nil(t, err)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("failed to verify user")

		repoMock.EXPECT().GetUserByID(gomock.Any(), userID).
			Return(user_model.User{}, expectedErr)

		err := service.VerifyUser(context.Background(), userID)

		require.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return error when user is already verified", func(t *testing.T) {
		repoRes.IsVerified = true
		expectedErr := util_error.NewBadRequest(errors.New("user is already verified"), "User is already verified")

		repoMock.EXPECT().GetUserByID(gomock.Any(), userID).
			Return(repoRes, nil)

		err := service.VerifyUser(context.Background(), userID)

		require.Error(t, err)
		assert.EqualError(t, err, expectedErr.Error())
	})
}
