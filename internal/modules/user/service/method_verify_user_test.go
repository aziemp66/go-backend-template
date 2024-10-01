package user_service

import (
	user_model "backend-template/internal/modules/user/model"
	mock_repository "backend-template/mock/repository"
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
	service := userService{userRepository: repoMock}

	email := "test@example.com"

	repoRes := user_model.User{
		ID:         "123",
		Email:      email,
		Password:   "secure_password",
		Name:       "Joe",
		Address:    "Sesame Street",
		IsVerified: true,
	}

	t.Run("should verify user successfully", func(t *testing.T) {
		repoMock.EXPECT().GetUserByEmail(gomock.Any(), email).Return(repoRes, nil)

		repoMock.EXPECT().VerifyUser(gomock.Any(), email).
			Return(nil)

		err := service.VerifyUser(context.Background(), email)

		require.NoError(t, err)
		assert.Nil(t, err)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("failed to verify user")

		repoMock.EXPECT().VerifyUser(gomock.Any(), email).
			Return(expectedErr)

		err := service.VerifyUser(context.Background(), email)

		require.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})

	t.Run("should return error when user is already verified", func(t *testing.T) {
		expectedErr := util_error.NewBadRequest(errors.New("user is already verified"), "User is already verified")

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), email).
			Return(repoRes, nil)

		repoMock.EXPECT().VerifyUser(gomock.Any(), email).Return(user_model.User{}, expectedErr)

		err := service.VerifyUser(context.Background(), email)

		require.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}
