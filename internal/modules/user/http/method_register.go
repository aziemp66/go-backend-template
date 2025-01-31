package user_http

import (
	user_model "backend-template/internal/modules/user/model"
	util_http "backend-template/util/http"

	"github.com/gin-gonic/gin"
)

func (userHttpHandler *userHttpHandler) Register(ctx *gin.Context) {
	var req user_model.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	userID, err := userHttpHandler.userService.Register(ctx.Request.Context(), req.Email, req.Password, req.Name, req.Address)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Registered User", user_model.IDResponse{
		ID: userID,
	})
}
