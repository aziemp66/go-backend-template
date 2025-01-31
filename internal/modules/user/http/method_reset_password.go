package user_http

import (
	user_model "backend-template/internal/modules/user/model"
	util_error "backend-template/util/error"
	util_http "backend-template/util/http"

	"github.com/gin-gonic/gin"
)

func (userHttpHandler *userHttpHandler) ResetPassword(ctx *gin.Context) {
	var req user_model.TokenRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(util_error.NewBadRequest(err, err.Error()))
		return
	}

	var bodyReq user_model.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		ctx.Error(err)
		return
	}

	err := userHttpHandler.userService.ResetPassword(ctx.Request.Context(), req.Token, bodyReq.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Reset Password", nil)
}
