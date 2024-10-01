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

	resToken := "expected_toked"
	resID := "1"

	t.Run("should register a new user", func(t *testing.T) {
		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).Return(user_model.User{}, sql.ErrNoRows)

		repoMock.EXPECT().CreateUser(gomock.Any(), reqEmail, reqPassword, reqName, reqAddress).
			Return(resID, nil)

		id, err := service.Register(context.Background(), reqEmail, reqPassword, reqName, reqAddress)

		jwtUtilMock.EXPECT().
			GenerateAuthToken(reqEmail, reqName, util_jwt.USER_ROLE, gomock.Any()).
			Return(resToken, nil)

		mailMock.EXPECT().SentMessage(gomock.Any()).Return(nil)

		require.NoError(t, err)
		assert.Equal(t, resID, id)
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

	t.Run("should return client error when email is already used", func(t *testing.T) {
		repoRes := user_model.User{
			ID:         resID,
			Email:      reqEmail,
			Password:   reqPassword,
			Name:       reqName,
			Address:    reqAddress,
			IsVerified: true,
		}
		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).Return(repoRes, nil)

		id, err := service.Register(context.Background(), reqEmail, reqPassword, reqName, reqAddress)

		expectedErr := util_error.NewBadRequest(fmt.Errorf("%s is already registered", reqEmail), fmt.Sprintf("Email %s is already used", reqEmail))

		assert.Error(t, err)
		assert.Empty(t, id)
		assert.EqualError(t, err, expectedErr.Message)
	})
}
