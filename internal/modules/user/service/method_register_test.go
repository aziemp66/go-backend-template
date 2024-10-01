package user_service

import (
	"context"
	"errors"
	"testing"

	mock_repository "backend-template/mock/repository"
	mock_util "backend-template/mock/util"
	util_jwt "backend-template/util/jwt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtUtilMock := mock_util.NewMockJWTManager(ctrl)
	mailMock := mock_util.NewMockMailManager(ctrl)
	service := userService{userRepository: repoMock}

	reqEmail := "test@example.com"
	reqPassword := "password123"
	reqName := "Test User"
	reqAddress := "123 Test St"
	reqID := "1"

	resToken := "expected_toked"

	t.Run("should register a new user", func(t *testing.T) {
		repoMock.EXPECT().CreateUser(gomock.Any(), reqEmail, reqPassword, reqName, reqAddress).
			Return(reqID, nil)

		id, err := service.Register(context.Background(), reqEmail, reqPassword, reqName, reqAddress)

		jwtUtilMock.EXPECT().
			GenerateAuthToken(reqEmail, reqName, util_jwt.USER_ROLE, gomock.Any()).
			Return(resToken, nil)

		mailMock.EXPECT().SentMessage(gomock.Any()).Return(nil)

		require.NoError(t, err)
		assert.Equal(t, reqID, id)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("registration error")

		repoMock.EXPECT().CreateUser(gomock.Any(), reqEmail, reqPassword, reqName, reqAddress).
			Return("", expectedErr)

		id, err := service.Register(context.Background(), reqEmail, reqPassword, reqName, reqAddress)

		require.Error(t, err)
		assert.Empty(t, id)
		assert.Equal(t, expectedErr, err)
	})
}
