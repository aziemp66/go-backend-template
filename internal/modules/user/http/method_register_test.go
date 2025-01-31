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

func TestUserHttpHandlerRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := mock_service.NewMockUserService(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	handler := NewUserHttpHandler(serviceMock, jwtMock)

	url := "/user/register"

	util_logger.InitLogger(gin.TestMode, "test", "./user_http_handler.log")

	app := util_http.NewHTTPServer(gin.TestMode)
	app.Use(
		util_http_middleware.TraceIdAssignmentMiddleware(),
		util_http_middleware.LogHandlerMiddleware(),
		util_http_middleware.ErrorHandlerMiddleware(),
	)
	app.POST(url, handler.Register)

	reqBody := user_model.CreateUserRequest{
		Email:    "john@example.com",
		Name:     "John",
		Address:  "Sesame Street 123",
		Password: "unsecure_password123",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	t.Run("should register successfully", func(t *testing.T) {
		userID := "123"
		serviceMock.EXPECT().Register(gomock.Any(), reqBody.Email, reqBody.Password, reqBody.Name, reqBody.Address).
			Return(userID, nil)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		resBody := util_http.Response{}
		err := json.NewDecoder(w.Body).Decode(&resBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.NotEmpty(t, resBody)
	})

	t.Run("should return error when service error", func(t *testing.T) {
		expError := util_error.NewBadRequest(nil, "email is already used")

		serviceMock.EXPECT().Register(gomock.Any(), reqBody.Email, reqBody.Password, reqBody.Name, reqBody.Address).Return("", expError)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		resBody := util_http.Response{}
		err := json.NewDecoder(w.Body).Decode(&resBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		assert.NotEmpty(t, resBody)
	})
}
