package user_service

import (
	"context"
	"errors"
	"testing"

	user_model "backend-template/internal/modules/user/model"
	mock_repository "backend-template/mock/repository"
	mock_util "backend-template/mock/util"
	util_jwt "backend-template/util/jwt"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	passwordManagerMock := mock_util.NewMockPasswordManager(ctrl)
	jwtManagerMock := mock_util.NewMockJWTManager(ctrl)
	service := userService{userRepository: repoMock}

	reqEmail := "test@example.com"
	reqPassword := "password123"

	t.Run("should login user successfully", func(t *testing.T) {
		resUser := user_model.User{
			ID:         "123",
			Email:      reqEmail,
			Password:   reqPassword,
			Name:       "John Doe",
			Address:    "Sesame Street",
			IsVerified: true,
		}
		resToken := "generated_token"

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).
			Return(resUser, nil)

		passwordManagerMock.EXPECT().CheckPasswordHash(reqPassword, resUser.Password).
			Return(nil)

		jwtManagerMock.EXPECT().GenerateAuthToken(resUser.ID, resUser.Name, util_jwt.USER_ROLE, gomock.Any()).
			Return(resToken, nil)

		resToken, err := service.Login(context.Background(), reqEmail, reqPassword)

		require.NoError(t, err)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("db error")

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).
			Return(user_model.GetUserResponse{}, expectedErr)

		_, err := service.Login(context.Background(), reqEmail, reqPassword)

		require.Error(t, err)
	})
}
