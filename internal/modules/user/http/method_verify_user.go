package user_http

import (
	user_model "backend-template/internal/modules/user/model"
	util_error "backend-template/util/error"
	util_http "backend-template/util/http"

	"github.com/gin-gonic/gin"
)

func (userHttpHandler *userHttpHandler) VerifyUser(ctx *gin.Context) {
	var req user_model.TokenRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(util_error.NewBadRequest(err, err.Error()))
		return
	}

	err := userHttpHandler.userService.VerifyUser(ctx, req.Token)
	if err != nil {
		ctx.Error(err)
		return
	}

	util_http.SendResponseJson(ctx, "Successfully Verifiy User", nil)
}
