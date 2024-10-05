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
	util_jwt "backend-template/util/jwt"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceForgotPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)
	mailMock := mock_util.NewMockMailManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock, mailMock)

	email := "test@example.com"
	repoRes := user_model.User{
		ID:         "1",
		Email:      email,
		Password:   "secured_password",
		Name:       "john",
		Address:    "sesame street",
		IsVerified: true,
	}

	token := "token123"

	t.Run("should send forgot password email", func(t *testing.T) {
		// We expect the repository to handle the logic for sending a reset password email
		repoMock.EXPECT().GetUserByEmail(gomock.Any(), email).Return(repoRes, nil)

		jwtMock.EXPECT().GenerateAuthToken(repoRes.ID, repoRes.Name, util_jwt.USER_ROLE, gomock.Any()).Return(token, nil)

		mailMock.EXPECT().SentResetPassword(token, repoRes.Email).Return(nil)

		err := service.ForgotPassword(context.Background(), email)
		require.NoError(t, err)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("error from db")

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), email).Return(user_model.User{}, expectedErr)

		err := service.ForgotPassword(context.Background(), email)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("should return not found when user is not found", func(t *testing.T) {
		repoMock.EXPECT().GetUserByEmail(gomock.Any(), email).Return(user_model.User{}, sql.ErrNoRows)

		expectedErr := util_error.NewNotFound(sql.ErrNoRows, fmt.Sprintf("User with the email of %s doesn't exist", email))

		err := service.ForgotPassword(context.Background(), email)
		require.EqualError(t, err, expectedErr.Error())
	})
}
