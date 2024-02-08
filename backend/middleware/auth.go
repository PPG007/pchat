package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	model_user "pchat/model/user"
	pb_common "pchat/pb/common"
	"pchat/utils"
)

var (
	noAuthPaths = []string{
		"/users/login",
		"/users/register",
		"/users/registerApplications",
	}

	unauthorizedTokenAvailablePaths = []string{
		"/users/validOTP",
	}
)

func auth(ctx *gin.Context) {
	fullPath := ctx.FullPath()
	if utils.StrInArray(fullPath, &noAuthPaths) {
		ctx.Next()
		return
	}
	token := utils.GetToken(ctx)
	userClaim, err := model_user.ValidToken(ctx, token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, pb_common.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	utils.SetUserId(ctx, userClaim.UserId)
	if !userClaim.IsAuthorized {
		if utils.StrInArray(ctx.FullPath(), &unauthorizedTokenAvailablePaths) {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, pb_common.EmptyResponse{})
		}
		return
	}
	ctx.Next()
}
