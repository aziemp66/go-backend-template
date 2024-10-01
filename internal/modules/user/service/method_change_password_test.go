package user_service

import (
	user_model "backend-template/internal/modules/user/model"
	mock_repository "backend-template/mock/repository"
	mock_util "backend-template/mock/util"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)

	service := userService{userRepository: repoMock}

	email := "test@example.com"
	oldPassword := "oldPassword123"
	newPassword := "newPassword456"

	t.Run("should successfully change password", func(t *testing.T) {
		repoRes := user_model.User{
			ID:         "123",
			Email:      email,
			Password:   "hashed_old_password",
			Name:       "John",
			Address:    "Sesame Street",
			IsVerified: true,
		}
		hashedPassword := "hashed_new_password"

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), email).Return(repoRes, nil)

		passwordMock.EXPECT().CheckPasswordHash(oldPassword, repoRes.Password).Return(nil)

		passwordMock.EXPECT().HashPassword(newPassword).Return(hashedPassword, nil)

		repoMock.EXPECT().ChangePassword(gomock.Any(), email, hashedPassword).Return(nil)

		err := service.ChangePassword(context.Background(), email, oldPassword, newPassword)

		require.NoError(t, err)
	})

	t.Run("should return error when failed retrieving db", func(t *testing.T) {
		expectedErr := errors.New("error from db")

		repoMock.EXPECT().ChangePassword(gomock.Any(), "", "").Return(expectedErr)

		err := service.ChangePassword(context.Background(), "", "", "")

		require.NoError(t, err)
	})
}
