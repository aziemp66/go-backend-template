package user_http

import (
	user_model "backend-template/internal/modules/user/model"
	mock_service "backend-template/mock/service"
	mock_util "backend-template/mock/util"
	util_error "backend-template/util/error"
	util_http "backend-template/util/http"
	util_http_middleware "backend-template/util/http/middleware"
	util_jwt "backend-template/util/jwt"
	util_logger "backend-template/util/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserHttpHandlerUpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := mock_service.NewMockUserService(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	handler := NewUserHttpHandler(serviceMock, jwtMock)

	url := "/user"

	util_logger.InitLogger(gin.TestMode, "test", "./user_http_handler.log")

	app := util_http.NewHTTPServer(gin.TestMode)
	app.Use(
		util_http_middleware.TraceIdAssignmentMiddleware(),
		util_http_middleware.LogHandlerMiddleware(),
		util_http_middleware.ErrorHandlerMiddleware(),
		util_http_middleware.JWTAuthentication(jwtMock),
		util_http_middleware.JWTAuthorization(util_jwt.USER_ROLE),
	)
	app.PUT(url, handler.UpdateProfile)

	reqBody := user_model.UpdateUserRequest{
		Name:    "John",
		Address: "SesameStreet",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	userToken := "token123"
	tokenBearer := fmt.Sprintf("BEARER %s", userToken)

	claims := &util_jwt.AuthClaims{
		ID:   uuid.NewString(),
		Name: "John",
		Role: util_jwt.USER_ROLE,
	}

	t.Run("should successfully update profile", func(t *testing.T) {
		jwtMock.EXPECT().VerifyAuthToken(userToken).Return(claims, nil)
		serviceMock.EXPECT().UpdateUser(gomock.Any(), claims.ID, reqBody.Name, reqBody.Address).Return(nil)

		req := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(reqBodyBytes))
		req.Header.Add(util_http.HEADER_AUTH, tokenBearer)
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		resBody := util_http.Response{}
		err := json.NewDecoder(w.Body).Decode(&resBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.NotEmpty(t, resBody)
	})

	t.Run("should return error when service error", func(t *testing.T) {
		jwtMock.EXPECT().VerifyAuthToken(userToken).Return(claims, nil)
		expectErr := util_error.NewBadRequest(nil, "Address Should Not be Empty")
		serviceMock.EXPECT().UpdateUser(gomock.Any(), claims.ID, reqBody.Name, reqBody.Address).Return(expectErr)

		req := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(reqBodyBytes))
		req.Header.Add(util_http.HEADER_AUTH, tokenBearer)
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		resBody := util_http.Response{}
		err := json.NewDecoder(w.Body).Decode(&resBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
		assert.NotEmpty(t, resBody)
	})
}
