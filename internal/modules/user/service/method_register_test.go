package user_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

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
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock)

	reqEmail := "test@example.com"
	reqPassword := "password123"
	reqName := "Test User"
	reqAddress := "123 Test St"

	hashedPassword := "secured_password"

	resID := "1"
	token := "abcd123"
	duration := 1 * time.Hour

	t.Run("should register a new user", func(t *testing.T) {
		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).Return(user_model.User{}, sql.ErrNoRows)

		passwordMock.EXPECT().PasswordValidation(reqPassword).Return(nil)

		passwordMock.EXPECT().HashPassword(reqPassword).Return(hashedPassword, nil)

		repoMock.EXPECT().CreateUser(gomock.Any(), reqEmail, hashedPassword, reqName, reqAddress).
			Return(resID, nil)

		jwtMock.EXPECT().GenerateAuthToken(resID, reqName, util_jwt.USER_ROLE, duration).Return(token, nil)

		id, err := service.Register(context.Background(), reqEmail, reqPassword, reqName, reqAddress)

		require.NoError(t, err)
		assert.Equal(t, resID, id)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("db error")

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).Return(user_model.User{}, expectedErr)

		id, err := service.Register(context.Background(), reqEmail, reqPassword, reqName, reqAddress)

		assert.Error(t, err)
		assert.Empty(t, id)
		assert.EqualError(t, err, expectedErr.Error())
	})

	t.Run("should return client error when email is already used", func(t *testing.T) {
		repoRes := user_model.User{
			ID:       resID,
			Email:    reqEmail,
			Password: reqPassword,
			Name:     reqName,
			Address:  reqAddress,
		}
		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).Return(repoRes, nil)

		id, err := service.Register(context.Background(), reqEmail, reqPassword, reqName, reqAddress)

		expectedErr := util_error.NewBadRequest(fmt.Errorf("%s is already registered", reqEmail), fmt.Sprintf("Email %s is already used", reqEmail))

		assert.Error(t, err)
		assert.Empty(t, id)
		assert.EqualError(t, err, expectedErr.Error())
	})
}
