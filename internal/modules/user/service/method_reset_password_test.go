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
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)
	mailMock := mock_util.NewMockMailManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock, mailMock)

	token := "some-reset-token"
	newPassword := "new_password123"

	expectedUser := user_model.User{
		ID:       "123",
		Email:    "john@example.com",
		Password: "old_password",
		Name:     "John",
		Address:  "Sesame Street",
	}
	expClaims := &util_jwt.AuthClaims{
		ID:   expectedUser.ID,
		Name: expectedUser.Name,
		Role: util_jwt.USER_ROLE,
	}

	t.Run("should reset password successfully", func(t *testing.T) {
		securedPassword := "secured_password"

		jwtMock.EXPECT().VerifyAuthToken(token).
			Return(expClaims, nil)

		repoMock.EXPECT().GetUserByID(gomock.Any(), expClaims.ID).
			Return(expectedUser, nil)

		passwordMock.EXPECT().PasswordValidation(newPassword).Return(nil)

		passwordMock.EXPECT().HashPassword(newPassword).Return(securedPassword, nil)

		repoMock.EXPECT().ChangePassword(gomock.Any(), expectedUser.Email, securedPassword).Return(nil)

		err := service.ResetPassword(context.Background(), token, newPassword)

		require.NoError(t, err)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("failed to reset password")

		jwtMock.EXPECT().VerifyAuthToken(gomock.Any()).Return(&util_jwt.AuthClaims{}, nil)

		repoMock.EXPECT().GetUserByID(gomock.Any(), "").Return(user_model.User{}, expectedErr)

		err := service.ResetPassword(context.Background(), "", "")

		assert.EqualError(t, err, expectedErr.Error())
	})
}
