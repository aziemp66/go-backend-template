package user_service

import (
	"context"
	"errors"
	"testing"

	user_model "backend-template/internal/modules/user/model"
	mock_repository "backend-template/mock/repository"
	mock_util "backend-template/mock/util"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// TODO: create unit test when user with the email is not found
func TestUserServiceForgotPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	mailMock := mock_util.NewMockMailManager(ctrl)
	service := userService{userRepository: repoMock}

	email := "test@example.com"
	repoRes := user_model.User{
		ID:         "1",
		Email:      email,
		Password:   "secured_password",
		Name:       "john",
		Address:    "sesame street",
		IsVerified: true,
	}

	t.Run("should send forgot password email", func(t *testing.T) {
		// We expect the repository to handle the logic for sending a reset password email
		repoMock.EXPECT().GetUserByEmail(gomock.Any(), email).Return(repoRes, nil)

		mailMock.EXPECT().SentMessage(gomock.Any()).Return(nil)

		err := service.ForgotPassword(context.Background(), email)
		require.NoError(t, err)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("error from db")

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), email).Return(expectedErr)

		err := service.ForgotPassword(context.Background(), email)
		require.EqualError(t, err, expectedErr.Error())
	})
}
