package user_http

import (
	user_model "backend-template/internal/modules/user/model"
	mock_service "backend-template/mock/service"
	mock_util "backend-template/mock/util"
	util_error "backend-template/util/error"
	util_http "backend-template/util/http"
	util_http_middleware "backend-template/util/http/middleware"
	util_logger "backend-template/util/logger"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserHttpHandlerResetPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := mock_service.NewMockUserService(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	handler := NewUserHttpHandler(serviceMock, jwtMock)

	url := "/user/reset-password/:token"
	token := "token123"
	urlWithToken := "/user/reset-password/" + token

	util_logger.InitLogger(gin.TestMode, "test", "./user_http_handler.log")

	app := util_http.NewHTTPServer(gin.TestMode)
	app.Use(
		util_http_middleware.TraceIdAssignmentMiddleware(),
		util_http_middleware.LogHandlerMiddleware(),
		util_http_middleware.ErrorHandlerMiddleware(),
	)
	app.POST(url, handler.ResetPassword)

	reqBody := user_model.ResetPasswordRequest{
		Password: "unsecure_password123",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	t.Run("should successfully reset password", func(t *testing.T) {
		serviceMock.EXPECT().ResetPassword(gomock.Any(), token, reqBody.Password).Return(nil)

		req := httptest.NewRequest(http.MethodPost, urlWithToken, bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		resBody := util_http.Response{}
		err := json.NewDecoder(w.Body).Decode(&resBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.NotEmpty(t, resBody)
	})

	t.Run("should return error when service error", func(t *testing.T) {
		expectErr := util_error.NewBadRequest(nil, "Password should be longer than 8 characters")
		serviceMock.EXPECT().ResetPassword(gomock.Any(), token, reqBody.Password).Return(expectErr)

		req := httptest.NewRequest(http.MethodPost, urlWithToken, bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		resBody := util_http.Response{}
		err := json.NewDecoder(w.Body).Decode(&resBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		assert.NotEmpty(t, resBody)
	})
}
