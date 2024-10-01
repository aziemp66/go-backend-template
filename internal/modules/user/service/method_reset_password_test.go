package user_service

import (
	"context"
	"errors"
	"testing"

	user_model "backend-template/internal/modules/user/model"
	mock_repository "backend-template/mock/repository"
	mock_util "backend-template/mock/util"
	util_jwt "backend-template/util/jwt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceResetPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtUtilMock := mock_util.NewMockJWTManager(ctrl)
	service := userService{userRepository: repoMock}

	token := "some-reset-token"
	email := "test@example.com"
	newPassword := "new_secure_password123"

	expectedUser := user_model.User{
		ID:       "123",
		Email:    email,
		Password: "old_password",
		Name:     "John",
		Address:  "Sesame Street",
	}
	expClaims := util_jwt.AuthClaims{
		ID:   expectedUser.ID,
		Name: expectedUser.Name,
		Role: util_jwt.USER_ROLE,
	}

	t.Run("should reset password successfully", func(t *testing.T) {
		// repoMock.EXPECT().ResetPassword(gomock.Any(), token, email, newPassword).
		// 	Return(nil)

		jwtUtilMock.EXPECT().VerifyAuthToken(token).
			Return(expClaims, nil)

		repoMock.EXPECT().GetUserByID(gomock.Any(), expClaims.ID).
			Return(expectedUser, nil)

		repoMock.EXPECT().ChangePassword(gomock.Any(), email, newPassword)

		err := service.ResetPassword(context.Background(), token, email, newPassword)

		require.NoError(t, err)
		assert.Nil(t, err)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("failed to reset password")

		repoMock.EXPECT().ChangePassword(gomock.Any(), email, newPassword).Return(expectedErr)

		err := service.ResetPassword(context.Background(), token, email, newPassword)

		require.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}
